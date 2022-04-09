package ports

import (
	"wallet_engine/internals/core/domain"
)

// RequestDTO declaring input DTO
type RequestDTO interface {
	domain.Wallet
}
