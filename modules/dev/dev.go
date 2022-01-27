package dev

import (
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/itsLuuke/go_tgbot/modules"
	"github.com/itsLuuke/go_tgbot/modules/utils/cmdHandler"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
)

var cnf = &modules.Config

func leave(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	args := ctx.Args()
	u := ctx.EffectiveUser
	if u.Id != cnf.OwnerId {
		return ext.EndGroups
	} else {
		if len(args) >= 1 {
			if chat.Type == "private" {
				_, err := b.SendMessage(chat.Id, "Please specify a chat id", nil)
				logging.HandleErr(err)
			} else {
				_, err := b.SendMessage(chat.Id, "I'm about to leave this chat\nClick the button below to verify action.", &gotgbot.SendMessageOpts{
					ReplyMarkup: gotgbot.InlineKeyboardMarkup{
						InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
							{
								{
									Text:         "Leave",
									CallbackData: "leave_chat(" + strconv.FormatInt(chat.Id, 10) + ")",
								},
							},
						},
					},
				},
				)
				logging.HandleErr(err)
			}
		} else {
			chatId, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				_, err := b.SendMessage(chat.Id, "Please specify a valid chat id", nil)
				logging.HandleErr(err)
			} else {
				_, err = b.LeaveChat(chatId)
				if err != nil {
					_, er := b.SendMessage(chat.Id, "Left Chat!", nil)
					logging.HandleErr(er)
				}
			}
		}
	}
	return nil
}

func leave_cb(b *gotgbot.Bot, ctx *ext.Context) error {
	u := ctx.CallbackQuery.From.Id
	cq := ctx.CallbackQuery
	if u != cnf.OwnerId {
		cq.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "This is not for u",
			ShowAlert: false,
		},
		)
	} else {
		chatIdRaw := cq.Data[len("leave_chat(") : len(cq.Data)-1]
		chatId, er := strconv.ParseInt(chatIdRaw, 10, 64)
		if er != nil {
			cq.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "An error occured" + er.Error(), ShowAlert: false})
		}
		_, _, err := ctx.EffectiveMessage.EditText(b, "Left Chat!", &gotgbot.EditMessageTextOpts{})
		logging.HandleErr(err)
		_, errr := b.LeaveChat(chatId)
		if errr != nil {
			cq.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "An error occured" + er.Error(), ShowAlert: false})
		}
	}
	return nil
}

func LoadDev(d *ext.Dispatcher) {
	defer logging.Info("Loaded `dev` module")
	d.AddHandler(cmdHandler.NewCommand("leave", leave))
	d.AddHandler(handlers.NewCallback(callbackquery.Prefix("leave_chat"), leave_cb))
}
