package print

import (
	"encoding/json"
	"fmt"
	"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetTimestamp() string {
	now := time.Now().UTC()
	sec := now.UnixMilli()
	return fmt.Sprint(sec)
}
