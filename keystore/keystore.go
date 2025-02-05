package keystore

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

// EthereumMessageHashPrefix is a Geth-originating message prefix that seeks to
// prevent arbitrary message data to be representable as a valid Ethereum transaction
// For more information, see: https://github.com/ethereum/go-ethereum/issues/3731
const EthereumMessageHashPrefix = "\x19Ethereum Signed Message:\n32"

var ErrKeyStoreLocked = errors.New("keystore is locked (HINT: did you forget to call keystore.Unlock?)")

//go:generate mockery --name KeyStore --output ../internal/mocks/ --case=underscore
type KeyStore interface {
	Unlock(password string) error
	Accounts() []accounts.Account
	Wallets() []accounts.Wallet
	HasAccounts() bool
	HasAccountWithAddress(common.Address) bool
	NewAccount() (accounts.Account, error)
	Import(keyJSON []byte, oldPassword string) (accounts.Account, error)
	Export(address common.Address, newPassword string) ([]byte, error)
	Delete(address common.Address) error
	GetAccounts() []accounts.Account
	GetAccountByAddress(common.Address) (accounts.Account, error)

	SignTx(account accounts.Account, tx *ethTypes.Transaction, chainID *big.Int) (*ethTypes.Transaction, error)
}

// keyStore manages a key storage directory on disk.
type keyStore struct {
	*keystore.KeyStore
	password     string
	scryptParams ScryptParams
}

// NewKeyStore creates a keystore for the given directory.
func NewKeyStore(keyDir string, scryptParams ScryptParams) KeyStore {
	return &keyStore{
		keystore.NewKeyStore(keyDir, scryptParams.N, scryptParams.P),
		"",
		scryptParams,
	}
}

// NewInsecureKeyStore creates an *INSECURE* keystore for the given directory.
// NOTE: Should only be used for testing!
func NewInsecureKeyStore(keyDir string) KeyStore {
	return NewKeyStore(keyDir, FastScryptParams)
}

// HasAccounts returns true if there are accounts located at the keystore
// directory.
func (ks *keyStore) HasAccounts() bool {
	return len(ks.Accounts()) > 0
}

// Unlock uses the given password to try to unlock accounts located in the
// keystore directory.
func (ks *keyStore) Unlock(password string) error {
	var merr error
	for _, account := range ks.Accounts() {
		err := ks.KeyStore.Unlock(account, password)
		if err != nil {
			merr = multierr.Combine(merr, fmt.Errorf("invalid password for account %s", account.Address.Hex()), err)
		}
	}
	ks.password = password
	return merr
}

// NewAccount adds an account to the keystore
func (ks *keyStore) NewAccount() (accounts.Account, error) {
	if ks.password == "" {
		return accounts.Account{}, ErrKeyStoreLocked
	}

	acct, err := ks.KeyStore.NewAccount(ks.password)
	if err != nil {
		return accounts.Account{}, err
	}

	err = ks.KeyStore.Unlock(acct, ks.password)
	return acct, err
}

// SignTx uses the unlocked account to sign the given transaction.
func (ks *keyStore) SignTx(account accounts.Account, tx *ethTypes.Transaction, chainID *big.Int) (*ethTypes.Transaction, error) {
	return ks.KeyStore.SignTx(account, tx, chainID)
}

// GetAccounts returns all accounts
func (ks *keyStore) GetAccounts() []accounts.Account {
	return ks.Accounts()
}

func (ks *keyStore) HasAccountWithAddress(address common.Address) bool {
	for _, acct := range ks.Accounts() {
		if acct.Address == address {
			return true
		}
	}

	return false
}

// GetAccountByAddress returns the account matching the address provided, or an error if it is missing
func (ks *keyStore) GetAccountByAddress(address common.Address) (accounts.Account, error) {
	for _, account := range ks.Accounts() {
		if account.Address == address {
			return account, nil
		}
	}

	return accounts.Account{}, errors.New("no account found with that address")
}

func (ks *keyStore) Import(keyJSON []byte, oldPassword string) (accounts.Account, error) {
	if ks.password == "" {
		return accounts.Account{}, ErrKeyStoreLocked
	}

	acct, err := ks.KeyStore.Import(keyJSON, oldPassword, ks.password)
	if err != nil {
		return accounts.Account{}, errors.Wrap(err, "could not import ETH key")
	}

	err = ks.KeyStore.Unlock(acct, ks.password)
	return acct, err
}

func (ks *keyStore) Export(address common.Address, newPassword string) ([]byte, error) {
	if ks.password == "" {
		return nil, ErrKeyStoreLocked
	}

	acct, err := ks.GetAccountByAddress(address)
	if err != nil {
		return nil, errors.Wrap(err, "could not export ETH key")
	}

	return ks.KeyStore.Export(acct, ks.password, newPassword)
}

func (ks *keyStore) Delete(address common.Address) error {
	if ks.password == "" {
		return ErrKeyStoreLocked
	}

	acct, err := ks.GetAccountByAddress(address)
	if err != nil {
		return errors.Wrap(err, "could not delete ETH key")
	}

	return ks.KeyStore.Delete(acct, ks.password)
}
