package cmd

import (
	"fmt"
	"os"
)

func GetAddressFromEnv() string {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("HTTP_PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8000"
	}
	return fmt.Sprintf("%s:%s", host, port)
}
