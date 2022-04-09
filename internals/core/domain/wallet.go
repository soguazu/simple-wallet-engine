package domain

import (
	"github.com/satori/go.uuid"
)

// State defines the state of the wallet
type State string

const (
	// ACTIVE state for an active wallet
	ACTIVE State = "active"
	// INACTIVE state for an inactive wallet
	INACTIVE = "inactive"
)

// Wallet model
type Wallet struct {
	Base
	Owner   uuid.UUID `json:"owner," gorm:"not null;index"`
	Balance int64     `json:"balance" gorm:"not null"`
	Status  State     `json:"status" gorm:"index"`
}
