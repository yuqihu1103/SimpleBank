package util

// all supported currency
const (
	USD = "USD" // United States Dollar
	EUR = "EUR" // Euro
	JPY = "JPY" // Japanese Yen
	GBP = "GBP" // British Pound Sterling
	AUD = "AUD" // Australian Dollar
	CAD = "CAD" // Canadian Dollar
	CHF = "CHF" // Swiss Franc
	CNY = "CNY" // Chinese Yuan
	SEK = "SEK" // Swedish Krona
	NZD = "NZD" // New Zealand Dollar
	NOK = "NOK" // Norwegian Krone
	SGD = "SGD" // Singapore Dollar
	KRW = "KRW" // South Korean Won
	INR = "INR" // Indian Rupee
	BRL = "BRL" // Brazilian Real
	RMB = "RMB" // Chinese RenMinBi
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, GBP, CNY:
		return true
	}
	return false
}
