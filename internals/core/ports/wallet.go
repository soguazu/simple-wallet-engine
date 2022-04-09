package ports

import (
	"github.com/gin-gonic/gin"
	"wallet_engine/internals/common"
	"wallet_engine/internals/core/domain"
)

// IWalletService defines the interface for a wallet service
type IWalletService interface {
	GetWalletByID(id string) (*domain.Wallet, error)
	CreateWallet(wallet *domain.Wallet) error
	UpdateWallet(params common.GetByIDRequest, state common.UpdateWalletRequest) (*domain.Wallet, error)
	DeleteWallet(id string) error
}

// IWalletHandler defines the interface for company handler
type IWalletHandler interface {
	GetWalletByID(c *gin.Context)
	CreateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
	UpdateWallet(c *gin.Context)
}
