package common

import (
	uuid "github.com/satori/go.uuid"

	"wallet_engine/internals/core/domain"
)

// CreateWalletRequest DTO to create wallet
type CreateWalletRequest struct {
	Status string `json:"status" binding:"required"`
}

// CreateTransactionRequest DTO to create transaction
type CreateTransactionRequest struct {
	TransactionType string `json:"transaction_type" binding:"required"`
	Purpose         string `json:"purpose" binding:"required"`
	Amount          int64  `json:"amount" binding:"required"`
	AccountID       string `json:"account_id" binding:"required"`
}

// GetTransactionResponse DTO to create transaction
type GetTransactionResponse struct {
	TransactionType string `json:"transaction_type"`
	Purpose         string `json:"purpose"`
	Amount          int64  `json:"amount"`
	AccountID       string `json:"account_id"`
	BalanceBefore   int64  `json:"balance_before"`
	BalanceAfter    int64  `json:"balance_after"`
}

// CreateWalletResponse DTO return wallet
type CreateWalletResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    domain.Wallet `json:"data"`
}

// CreateTransactionResponse DTO return transaction
type CreateTransactionResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    domain.Transaction `json:"data"`
}

// GetWalletResponse DTO
type GetWalletResponse struct {
	ID        uuid.UUID    `json:"id" binding:"required"`
	Owner     uuid.UUID    `json:"owner" binding:"required"`
	Balance   int64        `json:"balance" binding:"required"`
	Status    domain.State `json:"status"`
	AccountID int32        `json:"account_id" binding:"required"`
}

// GetWalletByIDRequest DTO to get wallet by id
type GetWalletByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// UpdateWalletRequest DTO to update wallet
type UpdateWalletRequest struct {
	Status *string `json:"status,omitempty" form:"status"`
}

// Error struct
type Error struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Data to return generic data
type Data struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"error,omitempty"`
}

// Message struct
type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
