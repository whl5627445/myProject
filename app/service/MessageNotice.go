package service

import (
	"context"
	"encoding/json"
	"yssim-go/config"
)

func MessageNotice(mes interface{}) bool {
	mesJson, _ := json.Marshal(mes)
	err := config.R.LPush(context.Background(), config.USERNAME+"_"+"notification", mesJson)
	if err != nil {
		return false
	}
	return true
}
