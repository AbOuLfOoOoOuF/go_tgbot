package misc

import (
	"encoding/json"
	"fmt"
	"html"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/itsLuuke/go_tgbot/modules/utils/cmdHandler"
	"github.com/itsLuuke/go_tgbot/modules/utils/helpers"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
)

// todo: parse command
func getInfo(b *gotgbot.Bot, ctx *ext.Context) error {
	var u *gotgbot.Sender
	msg := ctx.EffectiveMessage
	if msg.ReplyToMessage != nil {
		u = msg.ReplyToMessage.GetSender()
	} else {
		u = ctx.EffectiveSender
	}
	var txt string = "<b>General Info</b>"
	txt += fmt.Sprintf("\n<b>ID:</b> <code>%d</code>", u.Id())
	txt += fmt.Sprintf("\n<b>Name:</b> %v", u.FirstName())
	if u.LastName() != "" {
		txt += fmt.Sprintf("\n<b>Last Name:</b> %v", u.LastName())
	}
	if u.Username() != "" {
		txt += fmt.Sprintf("\n<b>Username:</b> @%v", u.Username())
	}
	if u.IsUser() {
		txt += "\n<b>Link:</b> " + helpers.MentionUserHtml(u.Id(), u.FirstName())
	} else {
		txt += "\n<b>Link:</b> " + helpers.MentionChatHtml(u.Username(), u.FirstName())
	}
	if u.IsAutomaticForward {
		txt += "\n\nThis message is an automatic forward from the connected channel with ID: " + fmt.Sprintf("<code>%d</code>", u.Chat.Id)
	}
	_, err := msg.Reply(b, txt, &gotgbot.SendMessageOpts{
		ParseMode:                "html",
		DisableWebPagePreview:    true,
		ReplyToMessageId:         msg.MessageId,
		AllowSendingWithoutReply: true,
	})
	logging.HandleErr(err)
	return nil
}

func getId(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	u := ctx.EffectiveSender

	var txt string = "<b>Telegram IDs</b>"
	txt += fmt.Sprintf("\n<b>Chat:</b> <code>%d</code>", msg.Chat.Id)
	txt += fmt.Sprintf("\n<b>User:</b> <code>%d</code>", u.Id())

	if msg.ReplyToMessage != nil {
		if msg.ReplyToMessage.From != nil {
			txt += fmt.Sprintf("\n<b>Forward From:</b> <code>%d</code>", msg.ReplyToMessage.ForwardFrom.Id)
		}
	}

	_, err := msg.Reply(b, txt, &gotgbot.SendMessageOpts{
		ParseMode:                "html",
		ReplyToMessageId:         msg.MessageId,
		AllowSendingWithoutReply: true,
	})
	logging.HandleErr(err)
	return nil
}

func pingPong(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	startTime := time.Now().UnixNano()
	rmsg, err := msg.Reply(b, "Pong!", nil)
	endTime := time.Now().UnixNano()
	logging.HandleErr(err)
	text := fmt.Sprintf("<b>Pong!</b>\n<code>%d</code> ms", (endTime-startTime)/int64(time.Millisecond))

	_, _, err = rmsg.EditText(b, text, &gotgbot.EditMessageTextOpts{ParseMode: "html"})
	logging.HandleErr(err)
	return nil
}

func jsonify(b *gotgbot.Bot, ctx *ext.Context) error {
	m, err := json.MarshalIndent(ctx, "", "\t")
	logging.HandleErr(err)
	jsoned := fmt.Sprintf("<code>%s</code>", html.EscapeString(string(m)))
	_, err = ctx.EffectiveMessage.Reply(b, jsoned, &gotgbot.SendMessageOpts{
		ParseMode:                "html",
		ReplyToMessageId:         ctx.EffectiveMessage.MessageId,
		AllowSendingWithoutReply: true,
	})
	logging.HandleErr(err)
	return nil
}

func LoadMisc(d *ext.Dispatcher) {
	defer logging.Info("Loaded `misc` module")
	d.AddHandler(cmdHandler.NewCommand("info", getInfo))
	d.AddHandler(cmdHandler.NewCommand("id", getId))
	d.AddHandler(cmdHandler.NewCommand("ping", pingPong))
	d.AddHandler(cmdHandler.NewCommand("json", jsonify))
}
