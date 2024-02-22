package utils

import (
	"fmt"

	"InitiaRe-website/constant"
)

func GenerateUserKey(userId int) string {
	return fmt.Sprintf("%s: %d", constant.BasePrefix, userId)
}
