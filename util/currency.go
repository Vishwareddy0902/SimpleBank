package util

const (
	USD = "USD"
	EUR = "EUR"
	Rs  = "Rs"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, Rs:
		return true
	}
	return false
}
