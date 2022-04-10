package services

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"wallet_engine/internals/repositories"
	tx "wallet_engine/pkg/unit_of_work"

	"wallet_engine/internals/common"
	"wallet_engine/internals/core/domain"
	"wallet_engine/internals/core/ports"
)

type walletService struct {
	WalletRepository      repositories.Repository[domain.Wallet]
	TransactionRepository repositories.Repository[domain.Transaction]
	logger                *log.Logger
	db                    *gorm.DB
}

// NewWalletService function create a new instance for service
func NewWalletService(wr repositories.Repository[domain.Wallet], tr repositories.Repository[domain.Transaction], l *log.Logger, db *gorm.DB) ports.IWalletService {
	return &walletService{
		WalletRepository:      wr,
		TransactionRepository: tr,
		logger:                l,
		db:                    db,
	}
}

func (w *walletService) GetWalletByID(id string) (*domain.Wallet, error) {
	wallet, err := w.WalletRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (w *walletService) CreateWallet(wallet *domain.Wallet) error {
	err := w.WalletRepository.Persist(wallet)
	if err != nil {
		w.logger.Error(err)
		return err
	}
	return nil
}

func (w *walletService) DeleteWallet(id string) error {
	err := w.WalletRepository.Delete(id, domain.Wallet{})
	if err != nil {
		w.logger.Error(err)
		return err
	}
	return nil
}

func (w *walletService) UpdateWallet(params common.GetByIDRequest, body common.UpdateWalletRequest) (*domain.Wallet, error) {
	wallet, err := w.WalletRepository.GetByID(params.ID)
	if err != nil {
		w.logger.Error(err)
		return nil, err
	}
	if body.Status != nil {
		(*wallet).Status = domain.State(*body.Status)
	}

	err = w.WalletRepository.Update(wallet)

	if err != nil {
		w.logger.Error(err)
		return nil, err
	}
	return wallet, nil
}

func (w *walletService) CreateTransaction(params common.GetByIDRequest, body common.CreateTransactionRequest) (*domain.Transaction, error) {
	uw := tx.NewGormUnitOfWork(w.db)
	t, err := uw.Begin()

	defer func() {
		if err != nil {
			t.Rollback()
		}
	}()

	if err != nil {
		return nil, err
	}

	wallet, err := w.WalletRepository.GetByIDForUpdate(params.ID)

	if err != nil {
		return nil, err
	}

	transaction, err := w.ReturnTransaction(wallet, body)

	if err != nil {
		return nil, err
	}

	err = w.TransactionRepository.WithTx(t).Persist(transaction)

	if err != nil {
		return nil, err
	}

	(*wallet).Balance = transaction.BalanceAfter

	err = w.WalletRepository.Update(wallet)

	if err != nil {
		return nil, err
	}

	err = uw.Commit()

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (w *walletService) ReturnTransaction(wallet *domain.Wallet, transaction common.CreateTransactionRequest) (*domain.Transaction, error) {
	var total int64
	if transaction.TransactionType == "credit" {
		total = wallet.Balance + transaction.Amount
	} else {
		if wallet.Balance < transaction.Amount {
			return nil, errors.New("insufficient balance")
		}
		total = wallet.Balance - transaction.Amount
	}
	return &domain.Transaction{
		TransactionType: domain.TxnType(transaction.TransactionType),
		Purpose:         domain.PurposeType(transaction.Purpose),
		Amount:          transaction.Amount,
		BalanceBefore:   wallet.Balance,
		BalanceAfter:    total,
		AccountID:       wallet.AccountID,
	}, nil
}
