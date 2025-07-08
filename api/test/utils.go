package test

import (
	"fmt"
	"math/rand"
)

func generateRandomEmail() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	email := fmt.Sprintf("%s@echo.test", result)

	return email
}
