package transfers

type Amount struct {
	value    float64
	currency string
}

func NewAmount(value float64, currency string) *Amount {
	return &Amount{
		value:    value,
		currency: currency,
	}
}

// IsAmountLimitAllowed check whether amount to transfer does not got beyond the bank limits
func (a *Amount) IsAmountLimitAllowed() bool {
	return a.value < 10000
}
