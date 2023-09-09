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

func randomInt(min, max int64) int64 {
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

func randomOwner() string {
	return randomString(6)
}

func randomMoney() int64 {
	return randomInt(0, 1000000)
}

func randomCurrency() string {
	currencys := []string{"USD", "GBP", "RMB"}
	l := len(currencys)
	c := currencys[rand.Intn(l)]
	return c
}
