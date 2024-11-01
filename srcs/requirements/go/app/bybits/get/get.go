package get

import (
	"bot/bybits/print"
	"bot/bybits/sign"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequetJson(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, err
}

func GetPrice(symbol string, url_bybit string) Price {
	var curr Price
	url := fmt.Sprint(url_bybit, "/v2/public/tickers?symbol=", symbol)
	body, err := GetRequetJson(url)
	if err != nil {
		log.Println(err)
		return curr
	}
	jsonErr := json.Unmarshal(body, &curr)
	if jsonErr != nil {
		log.Println(jsonErr)
		return curr
	}
	return curr
}

func GetWallet(api string, api_secret string, url_bybit string) Wallet {
	var wall Wallet
	params := map[string]string{
		"api_key":   api,
		"coin":      "USDT",
		"timestamp": print.GetTimestamp(),
	}

	signature := sign.GetSigned(params, api_secret)
	url := fmt.Sprint(
		url_bybit,
		"/v2/private/wallet/balance?api_key=",
		api,
		"&coin=USDT",
		"&timestamp=",
		params["timestamp"],
		"&sign=",
		signature,
	)
	body, err := GetRequetJson(url)
	if err != nil {
		log.Println(err)
		return wall
	}
	jsonErr := json.Unmarshal(body, &wall)
	if jsonErr != nil {
		log.Println(jsonErr)
		return wall
	}
	return (wall)
}
