package util

// Constant for all supported currencies
const (
	EUR = "EUR"
	USD = "USD"
	GBP = "GBP"
	JPY = "JPY"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {

	case EUR, USD, GBP, JPY, CAD:
		return true
	}
	return false
}
