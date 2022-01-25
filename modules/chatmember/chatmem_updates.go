package chatmember

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/itsLuuke/go_tgbot/modules/utils/helpers"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
)

const (
	owner  string = "owner"
	adm    string = "administrator"
	mem    string = "member"
	restr  string = "restricted"
	left   string = "left"
	kicked string = "kicked"
)

func chatMemUpdated(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	oldStat := ctx.ChatMember.OldChatMember.GetStatus()
	newStat := ctx.ChatMember.NewChatMember.GetStatus()

	byWhoM := ctx.ChatMember.From
	toWhoM := ctx.ChatMember.NewChatMember.GetUser()
	byWho := helpers.MentionUserHtml(byWhoM.Id, byWhoM.FirstName)
	toWho := helpers.MentionUserHtml(toWhoM.Id, toWhoM.FirstName)

	switch oldStat {
	case owner:
		switch newStat {
		// Owner can only leave or transfer ownership
		case adm:
			logChan(b, chat, toWho+" transferred the ownership to "+byWho)
		case left:
			logChan(b, chat, "Chat owner "+toWho+" left chat!")
		}
	case adm:
		switch newStat {
		case mem:
			logChan(b, chat, toWho+" has been demoted by "+byWho)
		case restr:
			logChan(b, chat, toWho+" has been demoted and muted by "+byWho)
		case left:
			logChan(b, chat, toWho+" left chat!")
		case kicked:
			logChan(b, chat, toWho+" has been demoted and removed by "+byWho)
		case owner:
			logChan(b, chat, toWho+" is now the owner!")
		}
	case mem:
		switch newStat {
		case owner:
			logChan(b, chat, toWho+" has been promoted and is now the owner!")
		case adm:
			logChan(b, chat, toWho+" has been promoted by "+byWho)
		case restr:
			logChan(b, chat, toWho+" has been muted by "+byWho)
		case left:
			logChan(b, chat, toWho+" Left chat!")
		case kicked:
			logChan(b, chat, toWho+" has been kicked by "+byWho)
		}
	case restr:
		switch newStat {
		case owner:
			logChan(b, chat, toWho+" has been promoted and is now the owner!")
		case adm:
			logChan(b, chat, toWho+" has been promoted and unmuted by "+byWho)
		case mem:
			logChan(b, chat, toWho+" has been unmuted by "+byWho)
		case left:
			logChan(b, chat, toWho+" left chat!")
		case kicked:
			logChan(b, chat, toWho+" has been kicked by "+byWho)
		}
	case left:
		switch newStat {
		case owner:
			logChan(b, chat, toWho+" has been added and is now the owner!")
		case adm:
			logChan(b, chat, toWho+" has been added as an administrator! by "+byWho)
		case mem:
			if byWhoM.Id == toWhoM.Id {
				logChan(b, chat, toWho+" joined chat!")
			} else {
				logChan(b, chat, toWho+" has been added by "+byWho)
			}
		case restr:
			logChan(b, chat, toWho+" has been restricted by "+byWho)
		case kicked:
			logChan(b, chat, toWho+" has been kicked by "+byWho)
		}
	case kicked:
		switch newStat {
		case owner:
			logChan(b, chat, toWho+" has been unbanned and is now the owner!")
		case adm:
			logChan(b, chat, toWho+" has been unbanned and promoted by "+byWho)
		case mem:
			logChan(b, chat, toWho+" has been unbanned and added by "+byWho)
		case restr: // meh
			logChan(b, chat, toWho+" has been restricted by "+byWho)
		case left:
			logChan(b, chat, toWho+" has been unbanned by "+byWho)
		}
	}
	return nil
}

func chatAdmUpdated(b *gotgbot.Bot, ctx *ext.Context) error {
	return nil // auto update cache on admin change
}

func logChan(b *gotgbot.Bot, chat *gotgbot.Chat, msg string) {

	// TODO: specific log levels
	// formatted := chat.Title + " [<code>" + strconv.FormatInt(chat.Id, 10) + "</code>]\n"
	// formatted += msg
	// formatted += "\n" + "Event Time Stamp: " + time.Now().Format(time.RFC3339)
	formatted := msg

	// for now just send to chat
	_, err := b.SendMessage(chat.Id, formatted, &gotgbot.SendMessageOpts{ParseMode: "html"})
	logging.HandleErr(err)
}

func LoadChatMemUpdates(d *ext.Dispatcher) {
	defer logging.Info("Loaded `chat member updates` module")
	d.AddHandlerToGroup(handlers.NewChatMember(func(u *gotgbot.ChatMemberUpdated) bool { return true }, chatMemUpdated), 10)
	d.AddHandlerToGroup(handlers.NewChatMember(func(u *gotgbot.ChatMemberUpdated) bool {
		if u.NewChatMember != u.OldChatMember {
			if u.OldChatMember.GetStatus() == adm {
				return true
			} else if u.NewChatMember.GetStatus() == adm {
				return true
			}
		}
		return false
	}, chatAdmUpdated), 20)
	d.AddHandlerToGroup(handlers.NewMyChatMember(func(u *gotgbot.ChatMemberUpdated) bool {
		if u.NewChatMember != u.OldChatMember {
			if u.OldChatMember.GetStatus() == adm {
				return true
			} else if u.NewChatMember.GetStatus() == adm {
				return true
			}
		}
		return false
	}, chatAdmUpdated), 30)
}
