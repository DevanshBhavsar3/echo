package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
)

var API_URL = "http://localhost:3001"

var randomEmail = generateRandomEmail()

func generateRandomEmail() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	randomEmail := fmt.Sprintf("%s@echo.test", result)
	return randomEmail
}

func sendRequest(t *testing.T, method string, url string, data interface{}, cookies []*http.Cookie) []byte {
	t.Helper()

	body, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}

	res, err := client.Do(req)
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

func getToken(t *testing.T, email string) string {
	t.Helper()

	url := fmt.Sprintf("%v/api/v1/auth/login", API_URL)

	user := types.LoginBody{
		Email:    email,
		Password: "test@123",
	}

	res := sendRequest(t, "POST", url, user, nil)

	var data types.AuthResponse
	err := json.Unmarshal(res, &data)
	if err != nil {
		t.Error(err)
	}

	return data.Token
}

func getWebsiteId(t *testing.T, tokenCookie *http.Cookie) string {
	t.Helper()

	url := fmt.Sprintf("%v/api/v1/website", API_URL)

	website := types.AddWebsiteBody{
		Url:       "http://echo.test.com",
		Frequency: "3m",
		Regions:   []string{"IND"},
	}

	res := sendRequest(t, "POST", url, website, []*http.Cookie{tokenCookie})

	var data types.AddWebsiteResponse
	err := json.Unmarshal(res, &data)
	if err != nil {
		t.Error(err)
	}

	return data.Id
}
