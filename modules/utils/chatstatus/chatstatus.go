package chatstatus

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
)

// todo: cache admins?
func IsAdmin(b *gotgbot.Bot, chat *gotgbot.Chat, user *gotgbot.User) bool {
	if chat.Type == "private" {
		return true
	}
	mem, err := chat.GetMember(b, user.Id)
	logging.HandleErr(err)
	if mem.MergeChatMember().Status == "creator" || mem.MergeChatMember().Status == "administrator" {
		return true
	}
	return false
}

func CanPromote(b *gotgbot.Bot, chat *gotgbot.Chat, user *gotgbot.User) bool {
	if chat.Type == "private" {
		return true
	}
	mem, err := chat.GetMember(b, user.Id)
	logging.HandleErr(err)
	return mem.MergeChatMember().CanPromoteMembers || mem.MergeChatMember().Status == "creator"
}
