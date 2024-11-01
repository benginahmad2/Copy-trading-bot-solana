package bot

import (
	"bot/data"
	"bot/mysql"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMsg(msg string, user string, api data.Env) error {
	url := fmt.Sprint(
		"https://api.telegram.org/bot",
		api.Api_telegram,
		"/sendMessage")
	params := map[string]string{
		"text":    msg,
		"chat_id": fmt.Sprint("@", user),
	}
	json_data, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
	} else {
		boby, reqErr := http.Post(url, "application/json", bytes.NewBuffer(json_data))
		if reqErr != nil {
			log.Println(reqErr)
		}
		log.Println(boby.Body)
	}
	return nil
}

func BotParseMsg(
	msg string,
	user string,
	api *data.Env,
	order *data.Bot,
	update tgbotapi.Update,
) error {
	if msg == "/help" {
		sender := fmt.Sprint("/add      api:api_secret\n",
			"/delete  api")
		msgs := tgbotapi.NewMessage(update.Message.Chat.ID, sender)
		order.Botapi.Send(msgs)
	} else if strings.Index(msg, "/add ") > -1 {
		if strings.Index(msg, ":") < 0 {
			msgs := tgbotapi.NewMessage(update.Message.Chat.ID, "Error try again \n/add api:api_secret")
			order.Botapi.Send(msgs)
		} else {
			var msgs tgbotapi.MessageConfig
			api_bybit := msg[strings.Index(msg, " ")+1 : strings.Index(msg, ":")]
			api_secret := msg[strings.Index(msg, ":")+1:]
			api.AddApi(api_bybit, api_secret)
			// add api to database
			if mysql.CheckApi("db.api", order.Db, api_bybit) == true {
				mysql.Insert(api_bybit, api_secret, "api", order.Db)
				msgs = tgbotapi.NewMessage(update.Message.Chat.ID, "Api add")
			} else {
				msgs = tgbotapi.NewMessage(update.Message.Chat.ID, "Api are already add")
			}
			order.Botapi.Send(msgs)
		}
	} else if strings.Index(msg, "/delete") > -1 {
		msg = msg[strings.Index(msg, " ")+1:]
		sender := api.Delette(msg)
		err := mysql.DbDelete("api", msg, order.Db)
		if err != nil {
			msgs := tgbotapi.NewMessage(update.Message.Chat.ID, sender)
			order.Botapi.Send(msgs)
		} else {
			msgs := tgbotapi.NewMessage(update.Message.Chat.ID, sender)
			order.Botapi.Send(msgs)
		}
	} else {
		SendMsg(msg, "trading_bybit_wax", *api)
	}
	return nil
}
