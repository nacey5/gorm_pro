package logic

import (
	"gorm_pro/chatroom-all/global"
	"strings"
)

// logic/sensitive.go
func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
	}

	return content
}
