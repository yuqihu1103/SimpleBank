package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "qwertyuioplkjhgfdsaxzcvbnm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000000)
}

func RandomMoneyPN() int64 {
	return RandomInt(-1000000, 1000000)
}

func RandomCurrency() string {
	currencys := []string{"USD", "GBP", "RMB"}
	l := len(currencys)
	c := currencys[rand.Intn(l)]
	return c
}
