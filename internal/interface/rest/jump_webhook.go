package rest

import (
	"encoding/json"
	"net/http"
)

type JumpWebhook struct {
}

func (self *JumpWebhook) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var requestBody Request

	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		return
	}
}
