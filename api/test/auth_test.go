package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	t.Run("Register user with correct data.", func(t *testing.T) {
		body := []byte(`{
			"name": "test user",
			"email": "testuser@test.com",
			"avatar": "https://google.com",
			"password": "test@123"
		}`)

		resp, err := http.Post("http://localhost:3000", "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var user store.User
		err = json.Unmarshal(resBody, &user)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Email, "testuser@test.com")
	})
}
