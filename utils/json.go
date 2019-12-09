package utils

import (
	"encoding/json"
	"log"
)

func ToJson(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

func FromJson(jsonStr string, toObj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), toObj)
}
