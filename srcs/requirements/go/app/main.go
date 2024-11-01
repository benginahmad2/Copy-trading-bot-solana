package main

import (
	"bot/bybits/bot"
	"bot/data"
	"bot/telegram" // Ensure this exists
	"bot/get"      // Ensure this exists
	"bot/post"     // Ensure this exists
	"bot/mysql"    // Ensure this exists
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)



// Main entry point of the application
func main() {
	var api data.Env
	var order data.Bot
	var trade data.Trades

	log.Print("Waiting for MySQL to initialize...")
	retryDatabaseConnection(&order, &api)

	// Load environment variables
	if err := data.LoadEnv(&api); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize bot with environment settings
	if err := order.NewBot(&api, false); err != nil {
		log.Fatalf("NewBot error: %v", err)
	}
	defer order.Db.Close()

	// Display available APIs
	api.ListApi()
	log.Print("API configuration loaded successfully")

	// Initialize Telegram Bot API
	order.Botapi
}




// run function to handle Telegram updates and execute trades
func run(updates tgbotapi.UpdatesChannel, order *data.Bot, api *data.Env, trade *data.Trades) {
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := update.Message.Text

			// Parse message for trade action
			bot.BotParseMsg(msg, update.Message.From.UserName, api, order, update)
			dataBybite, err := telegram.ParseMsg(msg, order.Debeug)
			if err != nil {
				log.Printf("Error Parsing Message: %v", err)
				continue
			}

			// Handle trade actions
			if dataBybite.Trade {
				handleTradeAction(api, dataBybite, order, trade)
			} else if dataBybite.Cancel {
				handleCancelAction(api, dataBybite, order, trade)
			}
		}
	}
}

// handleTradeAction - handles trade-related actions
func handleTradeAction(api *data.Env, dataBybite data.BybiteData, order *data.Bot, trade *data.Trades) {
	price := get.GetPrice(dataBybite.Currency, api.Url)
	if price.RetCode == 0 && price.Result[0].BidPrice != "" {
		for _, apis := range api.Api {
			if trade.Add(apis, dataBybite, price, api.Url) {
				post.PostIsoled(apis, dataBybite.Currency, trade, api.Url, order.Debeug)
				err := post.PostOrder(dataBybite.Currency, apis, trade, api.Url, order.Debeug)
				if err != nil {
					log.Printf("PostOrder error: %v", err)
					trade.Delete(dataBybite.Currency)
				} else {
					order.AddActive(dataBybite.Currency)
				}
			} else if order.Debeug {
				log.Printf("You already traded this Symbol")
			}
			if order.Debeug {
				trade.Print()
			}
		}
	} else {
		log.Printf("Symbol not found or invalid response")
	}
}

// handleCancelAction - handles cancellation actions
func handleCancelAction(api *data.Env, dataBybite data.BybiteData, order *data.Bot, trade *data.Trades) {
	for _, apis := range api.Api {
		err := post.CancelOrder(dataBybite.Currency, apis, trade, api.Url)
		if err != nil {
			log.Printf("CancelOrder error: %v", err)
		}

		trd := data.GetTrade(dataBybite.Currency, trade)
		if trd != nil {
			price := get.GetPrice(dataBybite.Currency, api.Url)
			sl := post.CancelBySl(price, trd)
			if sl != "" {
				err := post.ChangeLs(apis, dataBybite.Currency, sl, trd.Type, api.Url)
				if err != nil {
					log.Printf("ChangeLs error: %v", err)
				} else {
					log.Printf("Position cancelled successfully")
				}
			}
		}
		trade.Delete(dataBybite.Currency)
		order.Delete(dataBybite.Currency)
	}
}

