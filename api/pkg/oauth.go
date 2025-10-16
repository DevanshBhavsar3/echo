package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DevanshBhavsar3/echo/common/config"
	"github.com/DevanshBhavsar3/echo/common/db/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type ProviderConfig struct {
	oauth2.Config
	GetOAuthUser func(token *oauth2.Token) (*store.User, error)
}

type GoogleUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type GithubUser struct {
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type GithubEmail struct {
	Email string `json:"email"`
}

var OAuthConfig = map[string]*ProviderConfig{
	"google": {
		Config: oauth2.Config{
			RedirectURL:  "http://localhost:3001/api/v1/oauth/google/callback",
			ClientID:     config.Get("GOOGLE_CLIENT_ID"),
			ClientSecret: config.Get("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
		GetOAuthUser: func(token *oauth2.Token) (*store.User, error) {
			req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken, nil)
			if err != nil {
				return nil, err
			}

			userData, err := SendRequest(req)
			if err != nil {
				return nil, err
			}

			var user GoogleUser
			err = json.Unmarshal(userData, &user)
			if err != nil {
				return nil, err
			}

			return &store.User{
				Email: user.Email,
				Name:  user.Name,
				Image: user.Picture,
			}, nil
		},
	},
	"github": {
		Config: oauth2.Config{
			RedirectURL:  "http://localhost:3001/api/v1/oauth/github/callback",
			ClientID:     config.Get("GITHUB_CLIENT_ID"),
			ClientSecret: config.Get("GITHUB_CLIENT_SECRET"),
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		},
		GetOAuthUser: func(token *oauth2.Token) (*store.User, error) {
			req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
			if err != nil {
				return nil, err
			}

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

			userData, err := SendRequest(req)
			if err != nil {
				return nil, err
			}

			var user GithubUser
			err = json.Unmarshal(userData, &user)
			if err != nil {
				return nil, err
			}

			// Get the user email
			req, err = http.NewRequest("GET", "https://api.github.com/user/emails", nil)
			if err != nil {
				return nil, err
			}

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

			emailData, err := SendRequest(req)
			if err != nil {
				return nil, err
			}

			var emails []GithubEmail
			err = json.Unmarshal(emailData, &emails)
			if err != nil {
				return nil, err
			}

			return &store.User{
				Email: emails[0].Email,
				Name:  user.Name,
				Image: user.AvatarURL,
			}, nil
		},
	},
}

func GenerateRandomState() (string, error) {
	data := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
