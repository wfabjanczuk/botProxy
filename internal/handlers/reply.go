package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type incoming struct {
	Payload   payload `json:"payload"`
	SecretKey string  `json:"secret_key"`
}

type payload struct {
	ChatId string `json:"chat_id"`
}

func (a *app) Reply(w http.ResponseWriter, r *http.Request) {
	errPrefix := "replying to webhook request failed: "

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.writeClientError(w, errors.New(errPrefix+err.Error()), "Invalid request body.")
		return
	}

	in := &incoming{}
	err = json.Unmarshal(body, in)
	if err != nil {
		a.writeClientError(w, errors.New(errPrefix+err.Error()), "Invalid JSON.")
		return
	}

	if in.SecretKey != a.conf.SecretKey {
		a.writeClientError(w, errors.New(errPrefix+"invalid secret key"), "Invalid Secret.")
		return
	}

	messageFromBot := "Hi! I am bot " + a.botId
	err = a.client.SendEvent(in.Payload.ChatId, a.botId, messageFromBot)
	if err != nil {
		a.writeServerError(w, errors.New(errPrefix+err.Error()), "Automatic reply failed.")
		return
	}

	a.writeSuccess(w, "Automatic reply sent.")
}
