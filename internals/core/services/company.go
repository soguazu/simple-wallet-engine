package services

import (
	log "github.com/sirupsen/logrus"
	"wallet_engine/internals/repositories"

	"wallet_engine/internals/common"
	"wallet_engine/internals/core/domain"
	"wallet_engine/internals/core/ports"
)

type walletService struct {
	WalletRepository repositories.Repository[domain.Wallet]
	logger           *log.Logger
}

// NewWalletService function create a new instance for service
func NewWalletService(cr repositories.Repository[domain.Wallet], l *log.Logger) ports.IWalletService {
	return &walletService{
		WalletRepository: cr,
		logger:           l,
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
	err := w.WalletRepository.Delete(id)
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
		(*wallet).Status = *body.Status
	}

	if body.Balance != nil {
		(*wallet).Balance = *body.Balance
	}

	err = w.WalletRepository.Update(wallet)

	if err != nil {
		w.logger.Error(err)
		return nil, err
	}
	return wallet, nil
}
