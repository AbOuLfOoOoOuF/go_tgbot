package main

import (
	"net/http"

	m "github.com/itsLuuke/go_tgbot/modules"
	"github.com/itsLuuke/go_tgbot/modules/load"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
	"go.uber.org/zap"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var AllowedUpdates = []string{
	"message",
	"edited_message",
	"my_chat_member",
	"callback_query",
	"chat_member",
	// ! unused updates
	// "channel_post",
	// "edited_channel_post",
	// "inline_query",
	// "chosen_inline_result",
	// "shipping_query",
	// "pre_checkout_query",
	// "poll",
	// "poll_answer",
	// "chat_join_request",
}
var cnf = m.Env()

func main() {
	logger := logging.InitLogger()

	m.Env()

	BOT_TOKEN := cnf.Token

	b, err := gotgbot.NewBot(
		BOT_TOKEN,
		&gotgbot.BotOpts{
			Client:      http.Client{},
			GetTimeout:  gotgbot.DefaultGetTimeout,
			PostTimeout: gotgbot.DefaultPostTimeout,
		})
	logging.PanicErr(err) // Panic if there's an error

	u := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: zap.NewStdLog(logger),
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				logging.HandleErr(err)
				return ext.DispatcherActionNoop
			},
			Panic:       nil,
			ErrorLog:    zap.NewStdLog(logger),
			MaxRoutines: 0,
		},
	})

	d := u.Dispatcher
	load.LoadModules(d)

	err = u.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		Timeout:            0,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Offset:         0,
			Limit:          0,
			Timeout:        0,
			AllowedUpdates: AllowedUpdates,
		},
	})
	logging.PanicErr(err)

	startMsg := "[  BOT  ] Bot started"
	logging.Info(startMsg)
	sendStartMsg(b)

	u.Idle()
}

func sendStartMsg(b *gotgbot.Bot) {
	_, err := b.SendMessage(cnf.OwnerId, "Bot is up!", nil)
	logging.HandleErr(err)
}
