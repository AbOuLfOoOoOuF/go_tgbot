package helpers

import (
	"fmt"
	"html"
)

func MentionUserHtml(userId int64, name string) string {
	return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", userId, html.EscapeString(name))
}

func MentionChatHtml(Username string, Title string) string {
	return fmt.Sprintf("<a href=\"t.me/%v\">%v</a>", Username, html.EscapeString(Title))
}
