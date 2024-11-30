package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateConsignmentID() string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomPart := make([]byte, 6)
	for i := range randomPart {
		randomPart[i] = charset[r.Intn(len(charset))]
	}

	return timestamp + string(randomPart)
}
