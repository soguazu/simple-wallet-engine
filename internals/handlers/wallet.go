package handlers

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"wallet_engine/internals/common"
	"wallet_engine/internals/common/types"
	"wallet_engine/internals/core/domain"
	"wallet_engine/internals/core/ports"
	"wallet_engine/pkg/utils"
)

type walletHandler struct {
	WalletService ports.IWalletService
	logger        *log.Logger
	handlerName   string
}

var (
	result  utils.Result
	message types.Messages
)

// NewWalletHandler function creates a new instance for wallet handler
func NewWalletHandler(cs ports.IWalletService, l *log.Logger, n string) ports.IWalletHandler {
	return &walletHandler{
		WalletService: cs,
		logger:        l,
		handlerName:   n,
	}
}

// GetWalletByID godoc
// @Summary      Get a wallet
// @Description  get wallet by ID
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Success      200  {object}  common.GetWalletResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [get]
func (wh *walletHandler) GetWalletByID(c *gin.Context) {
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet, err := wh.WalletService.GetWalletByID(params.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			wh.logger.Error(err)
			c.JSON(http.StatusNotFound, result.ReturnErrorResult(err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		wh.logger.Error(err)
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(wallet, message.GetResponseMessage(wh.handlerName, types.OKAY)))
}

// CreateWallet godoc
// @Summary      Create wallet
// @Description  creates a wallet
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param wallet body common.CreateWalletRequest true "active or inactive"
// @Success      200  {object}  common.GetWalletResponse
// @Failure      400  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet [post]
func (wh *walletHandler) CreateWallet(c *gin.Context) {
	var body common.CreateWalletRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet := &domain.Wallet{
		Owner:     uuid.NewV4(),
		Status:    domain.State(body.Status),
		Balance:   0,
		AccountID: (&utils.Faker{}).RandomAccount(1000000000, 9999999999),
	}

	err := wh.WalletService.CreateWallet(wallet)
	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, result.ReturnSuccessResult(wallet, message.GetResponseMessage(wh.handlerName, types.CREATED)))
}

// DeleteWallet godoc
// @Summary      Delete a wallet by ID
// @Description  delete wallet by id
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [delete]
func (wh walletHandler) DeleteWallet(c *gin.Context) {
	var query common.GetByIDRequest
	if err := c.ShouldBindUri(&query); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	err := wh.WalletService.DeleteWallet(query.ID)
	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, result.ReturnSuccessMessage(types.DELETED))
}

// UpdateWallet godoc
// @Summary      Update a wallet by ID
// @Description  activate or deactivate wallet by id
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Param        status query   string  true  "active or inactive"
// @Success      200  {object}  common.GetWalletResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id}/activate [patch]
func (wh *walletHandler) UpdateWallet(c *gin.Context) {
	var query common.UpdateWalletRequest
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	wallet, err := wh.WalletService.UpdateWallet(params, query)
	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(wallet, message.GetResponseMessage(wh.handlerName, types.UPDATED)))
}

// TransactionWallet godoc
// @Summary      Transaction on a wallet by ID
// @Description  debit or credit wallet by id
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Wallet ID"
// @Param wallet body common.CreateTransactionRequest true "Create transaction"
// @Success      200  {object}  common.CreateTransactionResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [patch]
func (wh *walletHandler) TransactionWallet(c *gin.Context) {
	var body common.CreateTransactionRequest
	var params common.GetByIDRequest
	if err := c.ShouldBindUri(&params); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	transaction, err := wh.WalletService.CreateTransaction(params, body)
	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusBadRequest, result.ReturnErrorResult(err.Error()))
		return
	}

	c.JSON(http.StatusOK, result.ReturnSuccessResult(transaction, message.GetResponseMessage(wh.handlerName, types.CREATED_TRANSACTION)))
}
