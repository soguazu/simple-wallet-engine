package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"wallet_engine/internals/core/domain"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"wallet_engine/internals/common"
	"wallet_engine/internals/core/services"
	"wallet_engine/internals/repositories"
	datastore "wallet_engine/pkg/database"
	"wallet_engine/pkg/logger"
	"wallet_engine/pkg/utils"
)

var (
	db               = datastore.NewSqliteDatabase()
	DBConnection     = db.ConnectDB(filepath.Join("..", "..", "wallet.db"))
	logging          = logger.NewLogger(log.New()).MakeLogger(filepath.Join("..", "..", "logs", "info"), true)
	walletRepository = repositories.NewRepository[domain.Wallet](DBConnection)
	walletService    = services.NewWalletService(*walletRepository, logging)
	handler          = NewWalletHandler(walletService, logging, "Wallet")
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func createWallet(t *testing.T) *common.CreateWalletResponse {
	r := SetupRouter()
	r.POST("/v1/wallet", handler.CreateWallet)
	entity := common.CreateWalletRequest{
		Owner: (&utils.Faker{}).RandomUUID(),
	}

	jsonValue, _ := json.Marshal(entity)
	request, err := http.NewRequest("POST", "/v1/wallet", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error occurred decoding")
		return nil
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var wallet *common.CreateWalletResponse

	_ = json.Unmarshal([]byte(response.Body.String()), &wallet)

	require.Equal(t, http.StatusCreated, response.Code)

	return wallet
}

func TestWalletHandler_CreateWallet(t *testing.T) {
	wallet := createWallet(t)
	require.NotEmpty(t, wallet)
}

func TestWalletHandler_GetWalletByID(t *testing.T) {
	wallet := createWallet(t)

	r := SetupRouter()

	q := url.Values{}
	q.Add("id", wallet.Data.ID.String())

	r.GET("/v1/wallet/:id", handler.GetWalletByID)

	request, err := http.NewRequest("GET", "/v1/wallet/", strings.NewReader(q.Encode()))
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var resp common.GetWalletResponse

	err = json.Unmarshal([]byte(response.Body.String()), &resp)

	if err != nil {
		return
	}

	require.Equal(t, http.StatusOK, response.Code)
	require.Equal(t, wallet.Data.ID, resp.ID)
	require.Equal(t, wallet.Data.Owner, resp.Owner)
	require.Equal(t, wallet.Data.Balance, resp.Balance)
	require.Equal(t, wallet.Data.Status, resp.Status)
}

func TestWalletHandler_DeleteWallet(t *testing.T) {
	r := SetupRouter()
	wallet := createWallet(t)

	r.DELETE("/v1/wallet/:id", handler.DeleteWallet)

	endpoint := fmt.Sprintf("/v1/wallet/%v", wallet.Data.ID.String())

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusNoContent, response.Code)
}

func TestWalletHandler_UpdateWallet(t *testing.T) {
	r := SetupRouter()
	wallet := createWallet(t)

	r.PATCH("/v1/wallet/:id", handler.UpdateWallet)

	endpoint := fmt.Sprintf("/v1/wallet/%v", wallet.Data.ID.String())

	var status domain.State = "active"

	body := common.UpdateWalletRequest{
		Status: &status,
	}

	bytePayload, err := json.Marshal(body)
	if err != nil {
		return
	}

	request, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(bytePayload))

	if err != nil {
		return
	}

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	require.Equal(t, http.StatusOK, response.Code)
}
