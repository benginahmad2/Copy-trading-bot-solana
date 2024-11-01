package data

import (
	"bot/bybits/get"
	"bot/bybits/print"
	"bot/bybits/telegram"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Trade struct {
	Symbol      string   `json:"symbol"`
	Type        string   `json:"type"`
	Order       string   `json:"order"`
	SymbolPrice string   `json:"symbolPrice"`
	Wallet      string   `json:"wallet"`
	Price       string   `json:"price"`
	Entry       string   `json:"entry"`
	Leverage    string   `json:"leverage"`
	Tp1Order    string   `json:"tp_1Order"`
	Tp2Order    string   `json:"tp_2Order"`
	Tp3Order    string   `json:"tp_3Order"`
	Tp1         string   `json:"tp1"`
	Tp2         string   `json:"tp2"`
	Tp3         string   `json:"tp3"`
	Sl          string   `json:"Sl"`
	Id          []string `json:"id"`
	Active      []string `json:"active"`
}

type (
	Trades []Trade
)

type BybitApi struct {
	Api        string
	Api_secret string
}

type Env struct {
	Api          []BybitApi
	Api_telegram string
	Url          string
	BotName      string
}

type Bot struct {
	Active  []Start
	Debeug  bool
	Botapi  *tgbotapi.BotAPI
	Updates tgbotapi.UpdatesChannel
	Db      *sql.DB
}

type Start struct {
	Symbol string
	Active bool
}

func (t *Env) AddApi(api string, api_secret string) {
	check := false
	elem := BybitApi{
		Api:        api,
		Api_secret: api_secret,
	}
	for _, ls := range (*t).Api {
		if ls.Api == elem.Api {
			check = true
		}
	}
	if check == false {
		(*t).Api = append((*t).Api, elem)
	}
}

func (t *Env) Delette(api string) string {
	ret := false
	ls := (*t).Api
	var tmp []BybitApi

	for i := 0; i < len(ls); i++ {
		if ls[i].Api != api {
			tmp = append(tmp, ls[i])
		} else {
			ret = true
		}
	}
	(*t).Api = tmp
	if ret == false {
		return "Api not found cannot be deletted"
	}
	return "Api deletted"
}

func (t Env) ListApi() {
	for i := 0; i < len(t.Api); i++ {
		log.Println(print.PrettyPrint(t.Api[i]))
	}
}

func (t *Bot) NewBot(api *Env, debeug bool) error {
	elem := Bot{
		Active: nil,
		Debeug: debeug,
	}
	*t = elem
	return nil
}

func (t *Bot) CheckPositon(pos get.Position) {
	if len(pos.Result) > 0 {
		if pos.Result[0].EntryPrice > 0 || pos.Result[1].EntryPrice > 0 ||
			pos.Result[0].LiqPrice > 0 || pos.Result[1].LiqPrice > 0 ||
			pos.Result[0].BustPrice > 0 || pos.Result[1].BustPrice > 0 {
			log.Print("Entry PRice ok")
			for i := 0; i < len((*t).Active); i++ {
				if (*t).Active[i].Symbol == pos.Result[0].Symbol {
					(*t).Active[i].Active = true
					log.Print("Trade actif")
				} else {
					(*t).Active[i].Active = false
				}
			}
		}
	}
}

func (t Bot) GetActive() []string {
	var tmp []string
	for i := 0; i < len(t.Active); i++ {
		tmp = append(tmp, t.Active[i].Symbol)
	}
	return tmp
}

func (t *Bot) GetActiveSymbol(symbol string) bool {
	ret := false
	for _, ls := range (*t).Active {
		if ls.Symbol == symbol {
			ret = ls.Active
		}
	}
	return ret
}

func (t *Bot) AddActive(symbol string) {
	ls := (*t).Active
	elem := Start{
		Symbol: symbol,
		Active: false,
	}

	ls = append(ls, elem)
	(*t).Active = ls
}

func (t *Bot) Delete(symbol string) {
	var tmp []Start
	ls := (*t).Active

	for i := 0; i < len(ls); i++ {
		if symbol != ls[i].Symbol {
			tmp = append(tmp, ls[i])
		}
	}
	(*t).Active = tmp
}

func (t *Trades) SetId(symbol string, id string) {
	ls := *t

	for i := 0; i < len(ls); i++ {
		if ls[i].Symbol == symbol {
			ls[i].Id = append(ls[i].Id, id)
		}
	}
}

func (t *Trades) GetTrades() *Trades {
	return t
}

func (t *Trades) SetSl(symbol string, sl string) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].Symbol == symbol {
			(*t)[i].Sl = sl
		}
	}
}

func (t *Trades) GetSymbolOrder() []string {
	ls := *t
	var ret []string
	var check bool

	for i := 0; i < len(ls); i++ {
		check = true
		if ret == nil {
			ret = append(ret, ls[i].Symbol)
		} else {
			for j := 0; j < len(ret); j++ {
				if ret[j] == ls[i].Symbol {
					check = false
				}
			}
			if check {
				ret = append(ret, ls[i].Symbol)
			}
		}
	}
	return ret
}

func (t *Trades) CheckSymbol(symbol string) bool {
	ls := *t

	for i := 0; i < len(ls); i++ {
		if ls[i].Symbol == symbol {
			return false
		}
	}
	return true
}

func GetTrade(symbol string, t *Trades) *Trade {
	ls := *t

	for i := 0; i < len(ls); i++ {
		if ls[i].Symbol == symbol {
			return &ls[i]
		}
	}
	return nil
}

func (t *Trades) Add(api BybitApi, data telegram.Data, price get.Price, url_bybit string) bool {
	wallet := get.GetWallet(api.Api, api.Api_secret, url_bybit)
	available := wallet.Result.Usdt.AvailableBalance / 3
	prices, _ := strconv.ParseFloat(price.Result[0].BidPrice, 8)
	log.Println(print.PrettyPrint(available))
	elem := Trade{
		Symbol:      data.Currency,
		Type:        data.Type,
		Order:       data.Order,
		SymbolPrice: price.Result[0].BidPrice,
		Wallet:      fmt.Sprint(RoundFloat(available, 4)),
		Entry:       data.Entry,
		Leverage:    data.Level,
		Tp1Order:    RoundFloat((available*50/100)/prices, 4),
		Tp2Order:    RoundFloat((available*25/100)/prices, 4),
		Tp3Order:    RoundFloat((available*15/100)/prices, 4),
		Tp1:         data.Tp1,
		Tp2:         data.Tp2,
		Tp3:         data.Tp3,
		Sl:          data.Sl,
	}
	log.Printf("available")
	if t.CheckSymbol(data.Currency) == false {
		log.Printf("Trade actif Symbol: %s", data.Currency)
		return false
	}
	*t = append(*t, elem)
	return true
}

func (t *Trades) Delete(symbol string) bool {
	tmp := &Trades{}
	check := 0
	for i := 0; i < len(*t); i++ {
		if (*t)[i].Symbol != symbol {
			*tmp = append(*tmp, (*t)[i])
		} else {
			check = 1
		}
	}
	*t = *tmp
	if check == 1 {
		return false
	}
	return true
}

func (t *Trades) Print() {
	ls := *t

	for i := 0; i < len(ls); i++ {
		log.Println(print.PrettyPrint(ls[i]))
	}
}

func (t *Trades) GetLen() int {
	ls := *t

	return len(ls)
}

// get Trade
func (t *Trades) GetSymbol(index int) string {
	ls := *t
	for i := 0; i < len(ls); i++ {
		if i == index {
			return ls[i].Symbol
		}
	}
	return ""
}

func (t *Trades) GetType(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Type
	}
	return ""
}

func (t *Trades) GetOrder(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Order
	}
	return ""
}

func (t *Trades) GetSymbolPrice(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.SymbolPrice
	}
	return ""
}

func (t *Trades) GetWallet(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Wallet
	}
	return ""
}

func (t *Trades) GetPrice(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Price
	}
	return ""
}

func (t *Trades) GetEntry(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Entry
	}
	return ""
}

func (t *Trades) GetTp1Order(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Tp1Order
	}
	return ""
}

func (t *Trades) GetLeverage(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Leverage
	}
	return ""
}

func (t *Trades) GetTp2Order(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Tp2Order
	}
	return ""
}

func (t *Trades) GetTp3Order(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Tp3Order
	}
	return ""
}

func (t *Trades) GetTp1(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Tp1
	}
	return ""
}

func (t *Trades) GetTp2(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Tp2
	}
	return ""
}

func (t *Trades) GetTp3(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Tp3
	}
	return ""
}

func (t *Trades) GetSl(symbol string) string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Sl
	}
	return ""
}

func (t *Trades) GetId(symbol string) []string {
	ret := GetTrade(symbol, t)
	if ret != nil {
		return ret.Id
	}
	return nil
}

func RoundFloat(val float64, precision uint) string {
	ratio := math.Pow(10, float64(precision))
	ret := math.Round(val*ratio) / ratio
	rets := fmt.Sprint(ret)
	return rets
}

func GetEnv(env *Env) error {
	api := os.Getenv("API")
	if api == "" {
		return errors.New("Api not found")
	}
	api_secret := os.Getenv("API_SECRET")
	if api_secret == "" {
		return errors.New("Api_secret not found")
	}
	env.Api_telegram = os.Getenv("API_TELEGRAM")
	if env.Api_telegram == "" {
		return errors.New("Api_telegram not found")
	}
	env.Url = os.Getenv("URL")
	if env.Url == "" {
		return errors.New("Url not found")
	}
	env.AddApi(api, api_secret)
	return nil
}

func LoadEnv(env *Env) error {
	err := GetEnv(env)
	if err != nil {
		return err
	}
	return nil
}
