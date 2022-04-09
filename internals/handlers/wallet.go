package handlers

import (
	"errors"
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
// @Param wallet body common.CreateWalletRequest true "Add company"
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
		Owner:   body.Owner,
		Balance: 0,
		Status:  domain.ACTIVE,
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
// @Description  update wallet by id
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Company ID"
// @Param wallet body common.UpdateWalletRequest true "Update wallet"
// @Success      200  {object}  common.GetWalletResponse
// @Failure      400  {object}  common.Error
// @Failure      404  {object}  common.Error
// @Failure      500  {object}  common.Error
// @Router       /wallet/{id} [patch]
func (wh *walletHandler) UpdateWallet(c *gin.Context) {
	var body common.UpdateWalletRequest
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

	company, err := wh.WalletService.UpdateWallet(params, body)
	if err != nil {
		wh.logger.Error(err)
		c.JSON(http.StatusInternalServerError, result.ReturnErrorResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, result.ReturnSuccessResult(company, message.GetResponseMessage(wh.handlerName, types.UPDATED)))
}
