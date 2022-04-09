package common

import (
	uuid "github.com/satori/go.uuid"

	"wallet_engine/internals/core/domain"
)

// CreateWalletRequest DTO to create company
type CreateWalletRequest struct {
	Owner uuid.UUID `json:"owner" binding:"required"`
}

// CreateWalletResponse DTO return wallet
type CreateWalletResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    domain.Wallet `json:"data"`
}

// GetWalletResponse DTO
type GetWalletResponse struct {
	ID      uuid.UUID    `json:"id" binding:"required"`
	Owner   uuid.UUID    `json:"owner" binding:"required"`
	Balance int64        `json:"balance" binding:"required"`
	Status  domain.State `json:"status"`
}

// GetWalletByIDRequest DTO to get wallet by id
type GetWalletByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// UpdateWalletRequest DTO to update wallet
type UpdateWalletRequest struct {
	Status  *domain.State `json:"name,omitempty"`
	Balance *int64        `json:"balance,omitempty"`
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
