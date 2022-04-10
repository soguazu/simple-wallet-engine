package domain

// TxnType defines the transaction type
type TxnType string

// PurposeType defines the purpose type
type PurposeType string

const (
	// CREDIT a transaction type to credit's a wallet
	CREDIT TxnType = "credit"

	// DEBIT a transaction type to debit's a wallet
	DEBIT = "debit"
)

const (
	// DEPOSIT transaction purpose type
	DEPOSIT PurposeType = "deposit"

	// WITHDRAWAL transaction purpose type
	WITHDRAWAL = "withdrawal"

	// REVERSAL transaction purpose type
	REVERSAL = "reversal"
)

// Transaction model
type Transaction struct {
	Base
	TransactionType TxnType     `json:"transaction_type" gorm:"not null"`
	Purpose         PurposeType `json:"purpose" gorm:"not null;index"`
	Amount          int64       `json:"amount" gorm:"not null"`
	AccountID       int64       `json:"account_id" gorm:"not null;index"`
	BalanceBefore   int64       `json:"balance_before" gorm:"not null"`
	BalanceAfter    int64       `json:"balance_after" gorm:"not null"`
}
