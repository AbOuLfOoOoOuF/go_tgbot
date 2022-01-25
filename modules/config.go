package modules

import (
	"os"
	"strconv"
	"time"

	"github.com/itsLuuke/go_tgbot/modules/utils/logging"
	"github.com/joho/godotenv"
)

type Conf struct {
	BotStartTime int64
	Token        string
	OwnerId      int64
}

var StartTime = time.Now().Unix()
var Config Conf

func Env() Conf {
	c := Conf{}
	err := godotenv.Load()
	logging.PanicErr(err)

	// Bot start time
	c.BotStartTime = StartTime

	// Bot token
	var token_err bool
	c.Token, token_err = os.LookupEnv("TOKEN")
	if !token_err {
		logging.Panic("TOKEN not found, exiting")
	}

	// Owner id
	owner_id, er := os.LookupEnv("OWNER_ID")
	if !er {
		logging.Panic("OWNER_ID not found, exiting")
	}
	var ownerId_err error
	c.OwnerId, ownerId_err = strconv.ParseInt(owner_id, 10, 64)
	if ownerId_err != nil {
		logging.Panic("OWNER_ID is wrong, exiting")
	}
	Config = c
	return c
}
