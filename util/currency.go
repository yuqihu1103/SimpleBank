package util

// all supported currency
const (
	USD = "USD"
	CAD = "CAD"
	GBP = "GBP"
	CNY = "CNY"
	RMB = "RMB"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, GBP, CNY:
		return true
	}
	return false
}
