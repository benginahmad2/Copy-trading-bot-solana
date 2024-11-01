package post

import (
	"bot/bybits/get"
	"bot/bybits/print"
	"bot/bybits/sign"
	"bot/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func PostOrder(symbol string, api data.BybitApi, trade *data.Trades, url_bybit string, debug bool) error {
	params := map[string]interface{}{
		"api_key":          api.Api,
		"side":             trade.GetType(symbol),
		"symbol":           symbol,
		"order_type":       "Limit",
		"price":            trade.GetEntry(symbol),
		"time_in_force":    "GoodTillCancel",
		"reduce_only":      false,
		"close_on_trigger": false,
		"stop_loss":        trade.GetSl(symbol),
	}
	// tp1
	_, err := sendPost(params, trade.GetTp1(symbol), api, trade, trade.GetTp1Order(symbol), url_bybit, debug)
	if err != nil {
		return err
	}
	// tp2
	_, err = sendPost(params, trade.GetTp2(symbol), api, trade, trade.GetTp2Order(symbol), url_bybit, debug)
	if err != nil {
		return err
	}
	// tp3
	_, err = sendPost(params, trade.GetTp3(symbol), api, trade, trade.GetTp3Order(symbol), url_bybit, debug)
	if err != nil {
		return err
	}
	return nil
}

func sendPost(
	params map[string]interface{},
	tp string,
	api data.BybitApi,
	trade *data.Trades,
	order string,
	url_bybit string,
	debug bool,
) (*http.Response, error) {
	var res Post

	params["take_profit"] = tp
	params["qty"] = order
	params["timestamp"] = print.GetTimestamp()
	params["sign"] = sign.GetSignedinter(params, api.Api_secret)
	json_data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	if debug {
		println(print.PrettyPrint(params))
	}
	url := fmt.Sprint(url_bybit, "/private/linear/order/create")
	req, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(json_data),
	)
	if err != nil {
		return req, err
	}
	json.NewDecoder(req.Body).Decode(&res)
	if res.RetCode != 0 {
		return nil, errors.New(res.RetMsg)
	}
	log.Println(print.PrettyPrint(res))
	trade.SetId(params["symbol"].(string), res.Result.OrderID)
	delete(params, "sign")
	delete(params, "take_profit")
	delete(params, "qty")
	delete(params, "timestamp")
	return req, nil
}

func PostIsoled(api data.BybitApi, symbol string, trade *data.Trades, url_bybit string, debug bool) error {
	var isolated Isolated
	params := map[string]interface{}{
		"api_key":       api.Api,
		"symbol":        symbol,
		"is_isolated":   true,
		"buy_leverage":  10,
		"sell_leverage": 10,
		"timestamp":     print.GetTimestamp(),
	}
	params["sign"] = sign.GetSignedinter(params, api.Api_secret)
	json_data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	url := fmt.Sprint(url_bybit, "/private/linear/position/switch-isolated")
	req, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	json.NewDecoder(req.Body).Decode(&isolated)
	if debug {
		log.Printf("post PostIsoled")
		log.Println(print.PrettyPrint(isolated))
	}
	log.Printf("Isolated active: %d", params["buy_leverage"])
	return nil
}

func CancelOrder(symbol string, api data.BybitApi, trade *data.Trades, url_bybit string) error {
	params := map[string]string{
		"api_key": api.Api,
		"symbol":  symbol,
	}
	err := PostCancelOrder(params, api, url_bybit)
	if err != nil {
		return err
	}
	log.Printf("Cancel order success: %s", symbol)
	return nil
}

func PostCancelOrder(params map[string]string, api data.BybitApi, url_bybit string) error {
	var cancel PostCancel

	params["timestamp"] = print.GetTimestamp()
	params["sign"] = sign.GetSigned(params, api.Api_secret)
	json_data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	delete(params, "sign")
	url := fmt.Sprint(url_bybit, "/private/linear/order/cancel-all")
	req, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		return err
	}
	json.NewDecoder(req.Body).Decode(&cancel)
	log.Println(print.PrettyPrint(cancel))
	if cancel.RetCode != 0 {
		return errors.New(cancel.RetMsg)
	}
	return nil
}

func CancelBySl(price get.Price, trade *data.Trade) string {
	if trade.Type == "Buy" {
		log.Println(print.PrettyPrint(price))
		val, _ := strconv.ParseFloat(price.Result[0].BidPrice, 4)
		val = val - (val * 0.01)
		return fmt.Sprintf("%.4f", val)
	} else if trade.Type == "Sell" {
		val, _ := strconv.ParseFloat(price.Result[0].BidPrice, 8)
		val = val - (val * 0.01)
		return fmt.Sprintf("%.4f", val)
	}
	return ""
}

func ChangeLs(api data.BybitApi, symbol string, sl string, side string, url_bybit string) error {
	var stop StopLoss
	log.Println(symbol)
	log.Println(sl)
	log.Println(side)
	params := map[string]string{
		"api_key":   api.Api,
		"symbol":    symbol,
		"side":      side,
		"stop_loss": sl,
		"timestamp": print.GetTimestamp(),
	}
	params["sign"] = sign.GetSigned(params, api.Api_secret)
	json_data, err := json.Marshal(params)
	if err != nil {
		return err
	}
	url := fmt.Sprint(url_bybit, "/private/linear/position/trading-stop")
	req, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(json_data),
	)
	if err != nil {
		return err
	}
	json.NewDecoder(req.Body).Decode(&stop)
	log.Print("ChangeLs:")
	log.Println(print.PrettyPrint(stop))
	return nil
}
