package load

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/ratelimiter/ratelimiter"
	"github.com/itsLuuke/go_tgbot/modules"
	"github.com/itsLuuke/go_tgbot/modules/chatmember"
	"github.com/itsLuuke/go_tgbot/modules/start"
)

var cnf = &modules.Config

func LoadModules(d *ext.Dispatcher) {
	limiter := ratelimiter.NewLimiter(d, true, true)
	limiter.SetTriggerFuncs(limitedTrigger)
	limiter.AddExceptionID(cnf.OwnerId)
	limiter.Start()

	start.LoadStart(d)
	chatmember.LoadChatMemUpdates(d)

}

func limitedTrigger(b *gotgbot.Bot, ctx *ext.Context) error {
	var msg string
	if ctx.EffectiveSender.Chat != nil {
		msg = fmt.Sprintf("Channel [%v](t.me/%v) was spamming in [%v](t.me/c/%v) and is limited", ctx.EffectiveSender.Name(), ctx.EffectiveSender.Username(), ctx.EffectiveChat.Title, ctx.EffectiveChat.Id)
	} else {
		msg = fmt.Sprintf("User [%v](t.me/%v) was spamming in [%v](t.me/c/%v) and is limited", ctx.EffectiveSender.Name(), ctx.EffectiveSender.Username(), ctx.EffectiveChat.Title, ctx.EffectiveChat.Id)
	}
	b.SendMessage(cnf.OwnerId, msg, &gotgbot.SendMessageOpts{ParseMode: "markdown", DisableWebPagePreview: true})

	return nil
}
