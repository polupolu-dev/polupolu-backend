package validater

import (
	"regexp"
)

// UUID のバリデーション
func IsValidUUID(id string) bool {
	uuidRegex := regexp.MustCompile(`^[a-fA-F0-9-]{36}$`)
	return uuidRegex.MatchString(id)
}
