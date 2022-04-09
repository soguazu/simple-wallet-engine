package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"wallet_engine/internals/core/domain"
	"wallet_engine/internals/core/services"

	"wallet_engine/internals/handlers"
	"wallet_engine/internals/repositories"
	"wallet_engine/pkg/config"
	"wallet_engine/pkg/logger"
)

// Injection inject all dependencies
func Injection() {
	var logging *log.Logger

	if config.Instance.Env == "development" {
		logging = logger.NewLogger(log.New()).MakeLogger("logs/info", true)
		logging.Info("Log setup complete")
	} else {
		logging = logger.NewLogger(log.New()).Hook()
	}

	var (
		ginRoutes        = NewGinRouter(gin.Default())
		walletRepository = repositories.NewRepository[domain.Wallet](DBConnection)
		walletService    = services.NewWalletService(*walletRepository, logging)
		walletHandler    = handlers.NewWalletHandler(walletService, logging, "Wallet")
	)

	v1 := ginRoutes.GROUP("v1")
	wallet := v1.Group("/wallet")
	wallet.GET("/:id", walletHandler.GetWalletByID)
	wallet.POST("/", walletHandler.CreateWallet)
	wallet.DELETE("/:id", walletHandler.DeleteWallet)
	wallet.PATCH("/:id", walletHandler.UpdateWallet)

	err := ginRoutes.SERVE()

	if err != nil {
		return
	}

}
