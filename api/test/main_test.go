package test

import (
	"fmt"

	"github.com/DevanshBhavsar3/echo/api/config"
)

var API_URL string

func init() {
	port := config.GetEnv("PORT", "3000")

	API_URL = fmt.Sprintf("http://localhost:%v", port)
}
