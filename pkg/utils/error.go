package utils

import (
	"fmt"
	"strings"
)

func NewError(code string, msg string) error {
	return fmt.Errorf("%s:%s", code, msg)
}

func GetErrorMessage(err error) string {
	return strings.Split(err.Error(), ":")[1]
}
