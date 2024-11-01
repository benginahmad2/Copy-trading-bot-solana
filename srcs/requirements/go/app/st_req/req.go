package requet

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequetJson(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panic(err)
	}
	return body, err
}
