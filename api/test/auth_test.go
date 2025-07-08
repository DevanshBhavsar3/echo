package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/common/db/store"

	"github.com/stretchr/testify/assert"
)

var email = generateRandomEmail()

func TestRegister(t *testing.T) {
	registerUrl := fmt.Sprintf("%v/api/v1/auth/register", API_URL)

	t.Run("Register user with correct data.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		res := sendRequest(t, registerUrl, user)

		var data store.User
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, data.Email)
	})

	t.Run("Fails to register user with correct data.", func(t *testing.T) {
		body := types.RegisterUserBody{
			Name:   "test user",
			Email:  "testuser@test.com",
			Avatar: "https://google.com",
		}

		res := sendRequest(t, registerUrl, body)

		var data ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid data.", data.Error)
	})

	t.Run("Fails to register user with same email.", func(t *testing.T) {
		email := generateRandomEmail()

		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		res := sendRequest(t, registerUrl, user)

		var data store.User
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		res = sendRequest(t, registerUrl, user)

		var data1 ErrorResponse
		err = json.Unmarshal(res, &data1)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "User already exists.", data1.Error)
	})
}

func TestSignin(t *testing.T) {
	signinUrl := fmt.Sprintf("%v/api/v1/auth/signin", API_URL)

	t.Run("Signin user with correct data.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		res := sendRequest(t, signinUrl, user)

		var data store.User
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, data.Email)
	})

	t.Run("Fails to sign in user with incorrect password.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    email,
			Avatar:   "https://google.com",
			Password: "invalid@123",
		}

		res := sendRequest(t, signinUrl, user)

		var data ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid password.", data.Error)
	})

	t.Run("Fails to sign in with unknown email.", func(t *testing.T) {
		user := types.RegisterUserBody{
			Name:     "test user",
			Email:    "invalid@echo.test",
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		res := sendRequest(t, signinUrl, user)

		var data ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "User does not exists.", data.Error)
	})
}
