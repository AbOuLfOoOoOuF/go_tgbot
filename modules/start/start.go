package start

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/itsLuuke/go_tgbot/modules"
	"github.com/itsLuuke/go_tgbot/modules/utils/cmdHandler"
	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
)

var cnf = &modules.Config

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	chatId := ctx.EffectiveChat.Id

	upSec := cnf.BotStartTime
	upTime := time.Duration(time.Now().Unix()-upSec) * time.Second

	msg := "Hey!\nI'm alive since " + fmt.Sprintf("%v", upTime)
	_, err := bot.SendMessage(chatId, msg, nil)
	logging.HandleErr(err)
	return nil
}

func LoadStart(d *ext.Dispatcher) {
	defer logging.Info("Loaded `start` module")
	d.AddHandler(cmdHandler.NewCommand("start", Start))
}
