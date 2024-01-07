package generate

import (
	"math/rand"
	"time"
)

func RandomPassword(length int) string {
	const LOWER_ALPHA = "abcdefghijklmnopqrstuvwxyz"
	const UPPER_ALPHA = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const NUMERIC = "0123456789"
	const SYMBOLS = "!@#$%^&*()_+{}[];':\",./<>?"

	var str string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := []rune(LOWER_ALPHA + UPPER_ALPHA + NUMERIC + SYMBOLS)

	for i := 0; i < length; i++ {
		str += string(runes[r.Intn(len(runes))])
	}
	return str
}
