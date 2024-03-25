package auth

import (
	"os"
)

func GetKey(key string) string {
	return os.Getenv(key)
}
