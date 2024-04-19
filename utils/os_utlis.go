package utils

import (
	"errors"
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		errorText := fmt.Sprintf("%s ENV is not set", key)
		return "", errors.New(errorText)
	}
	return value, nil
}
