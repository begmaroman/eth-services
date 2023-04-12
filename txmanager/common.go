package txmanager

import (
	"bytes"
	"context"
	"math/big"
	"time"

	eskeystore "github.com/begmaroman/eth-services/keystore"
	"github.com/ethereum/go-ethereum"

	"github.com/begmaroman/eth-services/client"
	esStore "github.com/begmaroman/eth-services/store"
	"github.com/begmaroman/eth-services/store/models"
	"github.com/begmaroman/eth-services/types"
	"github.com/google/uuid"

	gethAccounts "github.com/ethereum/go-ethereum/accounts"
	gethCommon "github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

const (
	// maxEthNodeRequestTime is the worst case time we will wait for a response
	// from the eth node before we consider it to be an error
	maxEthNodeRequestTime = 30 * time.Second
)

func newAttempt(keyStore eskeystore.KeyStore, config *types.Config, tx *models.Tx, gasPrice *big.Int) (*models.TxAttempt, error) {
	attempt := models.TxAttempt{}
	account, err := keyStore.GetAccountByAddress(tx.FromAddress)
	if err != nil {
		return &attempt, errors.Wrapf(err, "error getting account %s for transaction %v", tx.FromAddress.String(), tx.ID)
	}

	transaction := gethTypes.NewTransaction(uint64(tx.Nonce), tx.ToAddress, tx.Value, tx.GasLimit, gasPrice, tx.EncodedPayload)
	hash, signedTxBytes, err := signTx(keyStore, account, transaction, config.ChainID)
	if err != nil {
		return &attempt, errors.Wrapf(err, "error using account %s to sign transaction %v", tx.FromAddress.String(), tx.ID)
	}

	attempt.ID = uuid.New()
	attempt.State = models.TxAttemptStateInProgress
	attempt.SignedRawTx = signedTxBytes
	attempt.TxID = tx.ID
	attempt.GasPrice = gasPrice
	attempt.Hash = hash
	attempt.BroadcastBeforeBlockNum = -1

	return &attempt, nil
}

func signTx(keyStore eskeystore.KeyStore, account gethAccounts.Account, tx *gethTypes.Transaction, chainID *big.Int) (gethCommon.Hash, []byte, error) {
	signedTx, err := keyStore.SignTx(account, tx, chainID)
	if err != nil {
		return gethCommon.Hash{}, nil, errors.Wrap(err, "signTx failed")
	}
	rlp := new(bytes.Buffer)
	if err := signedTx.EncodeRLP(rlp); err != nil {
		return gethCommon.Hash{}, nil, errors.Wrap(err, "signTx failed")
	}
	return signedTx.Hash(), rlp.Bytes(), nil

}

// sendTx broadcasts the tx to the ethereum network, writes any relevant data into the TxAttempt and
// returns an error (or nil) depending on the status.
func sendTx(ctx context.Context, txSender ethereum.TransactionSender, attempt *models.TxAttempt, logger types.Logger) *client.SendError {
	signedTx, err := attempt.GetSignedTx()
	if err != nil {
		return client.NewFatalSendError(err)
	}

	ctx, cancel := context.WithTimeout(ctx, maxEthNodeRequestTime)
	defer cancel()

	err = txSender.SendTransaction(ctx, signedTx)
	err = errors.WithStack(err)

	logger.Debugw("TxManager: Broadcasting transaction",
		"txAttemptID", attempt.ID,
		"txHash", signedTx.Hash(),
		"gasPriceWei", attempt.GasPrice.Int64(),
	)
	sendErr := client.NewSendError(err)
	if sendErr.IsTransactionAlreadyInMempool() {
		logger.Debugw("transaction already in mempool", "txHash", signedTx.Hash(), "nodeErr", sendErr.Error())
		return nil
	}
	return client.NewSendError(err)
}

// saveReplacementInProgressAttempt replace an old attempt with a new one.
func saveReplacementInProgressAttempt(store esStore.Store, tx *models.Tx, oldAttempt *models.TxAttempt, replacementAttempt *models.TxAttempt) error {
	errStr := "saveReplacementInProgressAttempt failed"
	if oldAttempt.State != models.TxAttemptStateInProgress || replacementAttempt.State != models.TxAttemptStateInProgress {
		return errors.Wrap(errors.New("expected attempts to be in_progress"), errStr)
	}
	if bytes.Equal(oldAttempt.ID[:], uuid.Nil[:]) {
		return errors.Wrap(errors.New("expected oldAttempt to have an ID"), errStr)
	}

	// Delete old attempt and add new attempt
	err := store.DeleteTxAttempt(oldAttempt.ID)
	if err != nil {
		return errors.Wrap(err, errStr)
	}
	err = store.PutTxAttempt(replacementAttempt)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	err = store.ReplaceAttempt(tx, oldAttempt, replacementAttempt)
	if err != nil {
		return errors.Wrap(err, errStr)
	}

	return nil
}

// BumpGas computes the next gas price to attempt as the largest of:
// - A configured percentage bump (GasBumpPercent) on top of the baseline price.
// - A configured fixed amount of Wei (GasBumpWei) on top of the baseline price.
// The baseline price is the maximum of the previous gas price attempt and TxManager's current gas price.
func BumpGas(config *types.Config, originalGasPrice *big.Int) (*big.Int, error) {
	baselinePrice := max(originalGasPrice, config.DefaultGasPrice)

	var priceByPercentage = new(big.Int)
	priceByPercentage.Mul(baselinePrice, big.NewInt(int64(100+config.GasBumpPercent)))
	priceByPercentage.Div(priceByPercentage, big.NewInt(100))

	var priceByIncrement = new(big.Int)
	priceByIncrement.Add(baselinePrice, config.GasBumpWei)

	bumpedGasPrice := max(priceByPercentage, priceByIncrement)
	if bumpedGasPrice.Cmp(config.MaxGasPrice) > 0 {
		return config.MaxGasPrice, errors.Errorf("bumped gas price of %s would exceed configured max gas price of %s (original price was %s)",
			bumpedGasPrice.String(), config.MaxGasPrice, originalGasPrice.String())
	} else if bumpedGasPrice.Cmp(originalGasPrice) == 0 {
		return bumpedGasPrice, errors.Errorf("bumped gas price of %s is equal to original gas price of %s."+
			" ACTION REQUIRED: This is a configuration error, you must increase either "+
			"GasBumpPercent or GasBumpWei", bumpedGasPrice.String(), originalGasPrice.String())
	}
	return bumpedGasPrice, nil
}
