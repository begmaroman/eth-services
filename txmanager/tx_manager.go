package txmanager

import (
	"math/big"

	"github.com/begmaroman/eth-services/client"
	esStore "github.com/begmaroman/eth-services/store"
	"github.com/begmaroman/eth-services/store/models"
	"github.com/begmaroman/eth-services/subscription"
	"github.com/begmaroman/eth-services/types"
	esTypes "github.com/begmaroman/eth-services/types"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type TxManager interface {
	Start() error

	RegisterAccount(address gethCommon.Address) error

	AddTx(
		from gethCommon.Address,
		to gethCommon.Address,
		value *big.Int,
		encodedPayload []byte,
		gasLimit uint64,
	) (uuid.UUID, error)

	GetTx(txID uuid.UUID) (*models.Tx, error)
	GetTxAttempt(attemptID uuid.UUID) (*models.TxAttempt, error)
	GetTxReceipt(receiptID uuid.UUID) (*models.TxReceipt, error)

	IsTxConfirmedAtOrBeforeBlockNumber(txID uuid.UUID, blockNumber int64) (bool, error)

	AddJob(txID uuid.UUID, metadata []byte) (uuid.UUID, error)

	MonitorJob(jobID uuid.UUID, handler JobHandler)

	DeleteJob(jobID uuid.UUID) error

	GetUnhandledJobIDs() ([]uuid.UUID, error)
}

type txManager struct {
	store  esStore.Store
	config *types.Config

	headTracker *subscription.HeadTracker
	broadcaster TxBroadcaster

	jobMonitor *jobMonitor
}

func NewTxManager(
	ethClient client.Client,
	store esStore.Store,
	keyStore client.KeyStoreInterface,
	config *esTypes.Config,
) (TxManager, error) {
	broadcaster := NewTxBroadcaster(ethClient, store, keyStore, config)
	confirmer := NewTxConfirmer(ethClient, store, keyStore, config)
	jobMonitor := newJobMonitor(store, config)
	headTracker := subscription.NewHeadTracker(
		ethClient,
		store,
		config,
		[]subscription.HeadTrackable{confirmer, jobMonitor},
	)

	return &txManager{
		store:  store,
		config: config,

		broadcaster: broadcaster,
		headTracker: headTracker,
		jobMonitor:  jobMonitor,
	}, nil
}

func (txm *txManager) Start() error {
	err := txm.headTracker.Start()
	if err != nil {
		return err
	}
	return txm.broadcaster.Start()
}

func (txm *txManager) RegisterAccount(address gethCommon.Address) error {
	return txm.broadcaster.RegisterAccount(address)
}

func (txm *txManager) AddTx(
	fromAddress gethCommon.Address,
	to gethCommon.Address,
	value *big.Int,
	encodedPayload []byte,
	gasLimit uint64,
) (uuid.UUID, error) {
	txID := uuid.New()
	err := txm.broadcaster.AddTx(txID, fromAddress, to, value, encodedPayload, gasLimit)
	if err != nil {
		return uuid.Nil, err
	}
	return txID, nil
}

func (txm *txManager) GetTx(txID uuid.UUID) (*models.Tx, error) {
	return txm.store.GetTx(txID)
}

func (txm *txManager) GetTxAttempt(attemptID uuid.UUID) (*models.TxAttempt, error) {
	return txm.store.GetTxAttempt(attemptID)
}

func (txm *txManager) GetTxReceipt(receiptID uuid.UUID) (*models.TxReceipt, error) {
	return txm.store.GetTxReceipt(receiptID)
}

func (txm *txManager) IsTxConfirmedAtOrBeforeBlockNumber(txID uuid.UUID, blockNumber int64) (bool, error) {
	return txm.store.IsTxConfirmedAtOrBeforeBlockNumber(txID, blockNumber)
}

func (txm *txManager) AddJob(txID uuid.UUID, metadata []byte) (uuid.UUID, error) {
	jobID := uuid.New()
	job := &models.Job{
		ID:       jobID,
		TxID:     txID,
		Metadata: metadata,
		State:    models.JobStateUnhandled,
	}
	err := txm.store.PutJob(job)
	if err != nil {
		return uuid.Nil, err
	}
	return jobID, nil
}

func (txm *txManager) DeleteJob(jobID uuid.UUID) error {
	return txm.store.DeleteJob(jobID)
}

func (txm *txManager) GetUnhandledJobIDs() ([]uuid.UUID, error) {
	return txm.store.GetUnhandledJobIDs()
}

func (txm *txManager) MonitorJob(jobID uuid.UUID, handler JobHandler) {
	m := txm.jobMonitor
	m.lock.Lock()
	m.jobs[jobID] = handler
	m.lock.Unlock()
}
