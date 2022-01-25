package helpers

import (
	"fmt"
	"html"
)

func MentionUserHtml(userId int64, name string) string {
	return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", userId, html.EscapeString(name))
}
