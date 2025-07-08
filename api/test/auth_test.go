package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/common/db/store"

	"github.com/stretchr/testify/assert"
)

var email = generateRandomEmail()

func TestRegister(t *testing.T) {
	t.Run("Register user with correct data.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		body, err := json.Marshal(user)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.Post("http://localhost:3000/api/v1/auth/register", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var res store.User
		err = json.Unmarshal(resBody, &res)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, user.Email)
	})

	t.Run("Fails to register user with correct data.", func(t *testing.T) {
		body := []byte(`{
			"name": "test user",
			"email": "testuser@test.com",
			"avatar": "https://google.com"
		}`)

		resp, err := http.Post("http://localhost:3000/api/v1/auth/register", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var errorRes ErrorResponse
		err = json.Unmarshal(resBody, &errorRes)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid data.", errorRes.Error)
	})

	t.Run("Fails to register user with same email.", func(t *testing.T) {
		email := generateRandomEmail()

		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		body, err := json.Marshal(user)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.Post("http://localhost:3000/api/v1/auth/register", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var res store.User
		err = json.Unmarshal(resBody, &res)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, user.Email)

		resp, err = http.Post("http://localhost:3000/api/v1/auth/register", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var errorRes ErrorResponse
		err = json.Unmarshal(resBody, &errorRes)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "User already exists.", errorRes.Error)
	})
}

func TestSignin(t *testing.T) {
	t.Run("Signin user with correct data.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		body, err := json.Marshal(user)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.Post("http://localhost:3000/api/v1/auth/signin", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var res store.User
		err = json.Unmarshal(resBody, &res)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, user.Email)
	})

	t.Run("Fails to sign in user with incorrect password.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "invalid@123",
		}

		body, err := json.Marshal(user)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.Post("http://localhost:3000/api/v1/auth/signin", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var errorRes ErrorResponse
		err = json.Unmarshal(resBody, &errorRes)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid password.", errorRes.Error)
	})

	t.Run("Fails to sign in with unknown email.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    "invalid@echo.test",
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		body, err := json.Marshal(user)
		if err != nil {
			t.Error(err)
		}

		resp, err := http.Post("http://localhost:3000/api/v1/auth/signin", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var errorRes ErrorResponse
		err = json.Unmarshal(resBody, &errorRes)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "User does not exists.", errorRes.Error)
	})
}
