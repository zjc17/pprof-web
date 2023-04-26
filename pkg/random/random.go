package random

import (
	"math/rand"
	"time"
)

func GenerateRandomDigitalWithUpperCaseLetterCode(length int) string {
	const charset = "0123456789" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return generateRandomCode(charset, length)
}

func generateRandomCode(charset string, length int) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}
