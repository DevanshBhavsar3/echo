package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/DevanshBhavsar3/echo/api/internal/types"
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateWebsite(t *testing.T) {
	url := fmt.Sprintf("%v/api/v1/website", API_URL)
	token := getToken(t, randomEmail)

	tokenCookie := &http.Cookie{
		Name:  "token",
		Value: token,
	}

	t.Run("Create website with correct data.", func(t *testing.T) {
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

		if !assert.NotEmpty(t, data.Id) {
			t.Error("expected website id")
		}
	})

	t.Run("Fails to create website with incorrect data.", func(t *testing.T) {
		website := types.AddWebsiteBody{
			Frequency: "3m",
			Regions:   []string{"IND"},
		}

		res := sendRequest(t, "POST", url, website, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid body.", data.Error)
	})

	t.Run("Fails to create website with incorrect freq.", func(t *testing.T) {
		website := types.AddWebsiteBody{
			Url:       "http://echo.test.com",
			Frequency: "30m",
			Regions:   []string{"IND"},
		}

		res := sendRequest(t, "POST", url, website, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid body.", data.Error)
	})

	t.Run("Fails to create website with incorrect region.", func(t *testing.T) {
		website := types.AddWebsiteBody{
			Url:       "http://echo.test.com",
			Frequency: "3m",
			Regions:   []string{"ARE"},
		}

		res := sendRequest(t, "POST", url, website, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid region provided.", data.Error)
	})
}

func TestGetWebsite(t *testing.T) {
	token := getToken(t, randomEmail)

	tokenCookie := &http.Cookie{
		Name:  "token",
		Value: token,
	}

	t.Run("Get website data with correct id.", func(t *testing.T) {
		websiteId := getWebsiteId(t, tokenCookie)

		url := fmt.Sprintf("%v/api/v1/website/%v", API_URL, websiteId)

		res := sendRequest(t, "GET", url, nil, []*http.Cookie{tokenCookie})

		var data store.Website
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, websiteId, data.ID)
	})

	t.Run("Fails to get website with incorrect id.", func(t *testing.T) {
		websiteId := uuid.New()

		url := fmt.Sprintf("%v/api/v1/website/%v", API_URL, websiteId)

		res := sendRequest(t, "GET", url, nil, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Website not found.", data.Error)
	})

	t.Run("Fails to get website with invalid uuid.", func(t *testing.T) {
		websiteId := "echo-test"

		url := fmt.Sprintf("%v/api/v1/website/%v", API_URL, websiteId)

		res := sendRequest(t, "GET", url, nil, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Invalid website id.", data.Error)
	})

	t.Run("Fails to get website from another user.", func(t *testing.T) {
		websiteId := getWebsiteId(t, tokenCookie)

		url := fmt.Sprintf("%v/api/v1/auth/register", API_URL)

		// Register 2nd user
		user2Email := generateRandomEmail()
		user2 := types.RegisterUserBody{
			Name:     "test user 2",
			Email:    user2Email,
			Avatar:   "https://google.com",
			Password: "test@123",
		}

		sendRequest(t, "POST", url, user2, nil)

		// Get token for 2nd user
		user2Token := getToken(t, user2Email)
		tokenCookie := &http.Cookie{
			Name:  "token",
			Value: user2Token,
		}

		url = fmt.Sprintf("%v/api/v1/website/%v", API_URL, websiteId)

		res := sendRequest(t, "GET", url, nil, []*http.Cookie{tokenCookie})

		var data types.ErrorResponse
		err := json.Unmarshal(res, &data)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "Website not found.", data.Error)
	})
}
