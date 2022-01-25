package admin

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/itsLuuke/go_tgbot/modules/utils/chatstatus"
	"github.com/itsLuuke/go_tgbot/modules/utils/cmdHandler"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
)

// todo: parse command
func promote(b *gotgbot.Bot, ctx *ext.Context) error {
	c := ctx.EffectiveChat
	u := ctx.EffectiveUser
	m := ctx.EffectiveMessage
	if m.ReplyToMessage == nil {
		_, err := m.Reply(b, "Reply to a user to promote him!", &gotgbot.SendMessageOpts{})
		logging.HandleErr(err)
		return ext.EndGroups
	}
	userId := m.ReplyToMessage.From.Id

	if c.Type == "private" {
		_, err := m.Reply(b, "This command is for group chats only.", &gotgbot.SendMessageOpts{})
		logging.HandleErr(err)
		return ext.EndGroups
	}

	if !chatstatus.CanPromote(b, c, u) {
		_, err := b.SendMessage(c.Id, "You don't have permission to promote users.", &gotgbot.SendMessageOpts{
			AllowSendingWithoutReply: true})
		logging.HandleErr(err)
		return ext.EndGroups
	}

	bot, _ := c.GetMember(b, b.Id)

	bl, err := b.PromoteChatMember(c.Id, userId, &gotgbot.PromoteChatMemberOpts{
		IsAnonymous:         false,
		CanManageChat:       bot.MergeChatMember().CanManageChat,
		CanPostMessages:     bot.MergeChatMember().CanPostMessages,
		CanEditMessages:     bot.MergeChatMember().CanEditMessages,
		CanDeleteMessages:   bot.MergeChatMember().CanDeleteMessages,
		CanManageVoiceChats: bot.MergeChatMember().CanManageVoiceChats,
		CanRestrictMembers:  bot.MergeChatMember().CanRestrictMembers,
		CanPromoteMembers:   false,
		CanChangeInfo:       bot.MergeChatMember().CanChangeInfo,
		CanInviteUsers:      bot.MergeChatMember().CanInviteUsers,
		CanPinMessages:      bot.MergeChatMember().CanPinMessages,
	})
	if !bl {
		_, err := m.Reply(b, "Failed to promote user.\n"+err.Error(), &gotgbot.SendMessageOpts{AllowSendingWithoutReply: true})
		logging.HandleErr(err)
	} else {
		_, err = m.Reply(b, "Promoted successfully.", &gotgbot.SendMessageOpts{AllowSendingWithoutReply: true})
		logging.HandleErr(err)
	}
	return nil
}
func demote(b *gotgbot.Bot, ctx *ext.Context) error {
	c := ctx.EffectiveChat
	u := ctx.EffectiveUser
	m := ctx.EffectiveMessage
	if m.ReplyToMessage == nil {
		_, err := m.Reply(b, "Reply to a user to demote him!", &gotgbot.SendMessageOpts{})
		logging.HandleErr(err)
	}
	userId := m.ReplyToMessage.From.Id

	if c.Type == "private" {
		_, err := m.Reply(b, "This command is for group chats only.", &gotgbot.SendMessageOpts{})
		logging.HandleErr(err)
	}

	if !chatstatus.CanPromote(b, c, u) {
		_, err := b.SendMessage(c.Id, "You don't have permission to demote users.", &gotgbot.SendMessageOpts{
			AllowSendingWithoutReply: true})
		logging.HandleErr(err)
	}

	bl, err := b.PromoteChatMember(c.Id, userId, &gotgbot.PromoteChatMemberOpts{
		IsAnonymous:         false,
		CanManageChat:       false,
		CanPostMessages:     false,
		CanEditMessages:     false,
		CanDeleteMessages:   false,
		CanManageVoiceChats: false,
		CanRestrictMembers:  false,
		CanPromoteMembers:   false,
		CanChangeInfo:       false,
		CanInviteUsers:      false,
		CanPinMessages:      false,
	})
	if !bl {
		_, err := m.Reply(b, "Failed to demote user.\n"+err.Error(), &gotgbot.SendMessageOpts{AllowSendingWithoutReply: true})
		logging.HandleErr(err)
	} else {
		_, err = m.Reply(b, "Demoted successfully.", &gotgbot.SendMessageOpts{AllowSendingWithoutReply: true})
		logging.HandleErr(err)
	}
	return nil
}

func LoadAdmin(d *ext.Dispatcher) {
	defer logging.Info("Loaded `admin` module")
	d.AddHandler(cmdHandler.NewCommand("promote", promote))
	d.AddHandler(cmdHandler.NewCommand("demote", demote))
}
