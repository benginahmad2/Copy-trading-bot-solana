package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"sort"
)

func GetSigned(params map[string]string, key string) string {
	keys := make([]string, len(params))
	i := 0
	_val := ""
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		_val += k + "=" + params[k] + "&"
	}
	_val = _val[0 : len(_val)-1]
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, _val)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetSignedinter(params map[string]interface{}, key string) string {
	keys := make([]string, len(params))
	i := 0
	_val := ""
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		if params[k] == true {
			_val += k + "=" + "true" + "&"
		} else if params[k] == false {
			_val += k + "=" + "false" + "&"
		} else {
			str := fmt.Sprintf("%v", params[k])
			_val += k + "=" + str + "&"
		}
	}
	_val = _val[0 : len(_val)-1]
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, _val)
	return fmt.Sprintf("%x", h.Sum(nil))
}
