package utils

import (
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/constant"
)

func GenerateUserKey(userId int) string {
	return fmt.Sprintf("%s: %d", constant.BasePrefix, userId)
}
