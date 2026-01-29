package utility

import (
	"math/rand"
	"time"
)

var seededOnce bool

func seedOnce() {
	if !seededOnce {
		rand.Seed(time.Now().UnixNano())
		seededOnce = true
	}
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const charset1 = "abcdefghijklmnopqrstuvwxyz" + "0123456789"

func StringWithCharset(length int, charset string) string {
	if length <= 0 || len(charset) == 0 {
		return ""
	}
	seedOnce()
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset1)
}

func StringAll(length int) string {
	return StringWithCharset(length, charset)
}
