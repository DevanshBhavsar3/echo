package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/common/db/store"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	url := fmt.Sprintf("%v/api/v1/auth/register", API_URL)

	t.Run("Register user with correct data.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    randomEmail,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		res := sendRequest(t, "POST", url, user, nil)

		var data types.AuthResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, data.User.Email)
	})

	t.Run("Fails to register user with incorrect data.", func(t *testing.T) {
		body := types.RegisterUserBody{
			Name:   "test user",
			Email:  "testuser@test.com",
			Avatar: "https://google.com",
		}

		res := sendRequest(t, "POST", url, body, nil)

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid data.", data.Error)
	})

	t.Run("Fails to register user with same email.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    randomEmail,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		res := sendRequest(t, "POST", url, user, nil)

		var data types.AuthResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		res = sendRequest(t, "POST", url, user, nil)

		var data1 types.ErrorResponse
		err = json.Unmarshal(res, &data1)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Email already used.", data1.Error)
	})
}

func TestLogin(t *testing.T) {
	url := fmt.Sprintf("%v/api/v1/auth/login", API_URL)

	t.Run("Login user with correct data.", func(t *testing.T) {
		user := types.LoginBody{
			Email:    randomEmail,
			Password: "test@123",
		}

		res := sendRequest(t, "POST", url, user, nil)

		var data types.AuthResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		if !assert.NotEmpty(t, data.Token) {
			t.Error("expected authentication token")
		}

		assert.Equal(t, user.Email, data.User.Email)
	})

	t.Run("Fails to sign in user with incorrect password.", func(t *testing.T) {
		user := types.LoginBody{
			Email:    randomEmail,
			Password: "invalid@123",
		}

		res := sendRequest(t, "POST", url, user, nil)

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid password.", data.Error)
	})

	t.Run("Fails to sign in with unknown email.", func(t *testing.T) {
		user := types.LoginBody{
			Email:    "invalid@echo.test",
			Password: "test@123",
		}

		res := sendRequest(t, "POST", url, user, nil)

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "User does not exists.", data.Error)
	})
}

func TestAuth(t *testing.T) {
	url := fmt.Sprintf("%v/api/v1/auth/user", API_URL)

	t.Run("Get user data with correct token.", func(t *testing.T) {
		token := getToken(t, randomEmail)

		tokenCookie := &http.Cookie{
			Name:  "token",
			Value: token,
		}

		res := sendRequest(t, "GET", url, nil, []*http.Cookie{tokenCookie})

		var data store.User
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		if !assert.NotEmpty(t, data.Email) {
			t.Error("expected user email")
		}
	})

	t.Run("Fails to get user data without token.", func(t *testing.T) {
		res := sendRequest(t, "GET", url, nil, nil)

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Token not provided.", data.Error)
	})

	t.Run("Fails to get user data with malformed token.", func(t *testing.T) {
		token := getToken(t, randomEmail)
		malformedToken := fmt.Sprintf("%ve", token[:len(token)-1])

		tokenCookie := &http.Cookie{
			Name:  "token",
			Value: malformedToken,
		}

		res := sendRequest(t, "GET", url, nil, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid token.", data.Error)
	})
}
