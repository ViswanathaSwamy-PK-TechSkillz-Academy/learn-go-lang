package transfers

import "time"

// Transfer contains attributes of a transfer entity
type Transfer struct {
	id        string
	date      time.Time
	amount    float64
	accountId int64
}

// NewTransfer is the only way to create a transfer entity
func NewTransfer(id string, date time.Time, amount float64, accountId int64) *Transfer {
	return &Transfer{
		id:        id,
		date:      date,
		amount:    amount,
		accountId: accountId,
	}
}

// IsAmountLimitAllowed check whether amount to transfer does not got beyond the bank limits
func (t *Transfer) IsAmountLimitAllowed() bool {
	if t.amount >= 10000 {
		return false
	}
	return true
}

// IsDateAllowed checks whether transfer is executed in the current day
func (t *Transfer) IsDateAllowed() bool {
	now := time.Now()
	if t.date.After(now) {
		return false
	}
	return true
}
