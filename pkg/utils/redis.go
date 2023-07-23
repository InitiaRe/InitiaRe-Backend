package utils

import (
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
)

func GenerateUserKey(userId int) string {
	return fmt.Sprintf("%s: %d", constants.BasePrefix, userId)
}
