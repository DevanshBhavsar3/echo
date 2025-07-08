package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"
)

var API_URL = "http://localhost:3000"

func generateRandomEmail() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	email := fmt.Sprintf("%s@echo.test", result)

	return email
}

func sendRequest(t *testing.T, url string, data interface{}) []byte {
	body, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	return resBody
}
