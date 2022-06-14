package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type incomingRequest struct {
	Payload   payload `json:"payload"`
	SecretKey string  `json:"secret_key"`
}

type payload struct {
	ChatId string `json:"chat_id"`
	Event  event  `json:"event"`
}

type event struct {
	Postback postback `json:"postback"`
}

type postback struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func (a *app) Reply(w http.ResponseWriter, r *http.Request) {
	errPrefix := "replying to webhook request failed: "

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.writeClientError(w, errors.New(errPrefix+err.Error()), "Invalid request body.")
		return
	}

	incoming := &incomingRequest{}
	err = json.Unmarshal(body, incoming)
	if err != nil {
		a.writeClientError(w, errors.New(errPrefix+err.Error()), "Invalid JSON.")
		return
	}

	if incoming.SecretKey != a.conf.SecretKey {
		a.writeClientError(w, errors.New(errPrefix+"invalid secret key"), "Invalid Secret.")
		return
	}

	pb := incoming.Payload.Event.Postback
	chatId := incoming.Payload.ChatId

	if pb.Id == "transfer" && pb.Value == "yes" {
		err = a.client.TransferChat(chatId, "agent", []string{a.conf.HumanId})
	} else {
		messageFromBot := "Hi! I am bot " + a.botId + ". Do you want to talk to a human?"
		err = a.client.SendEvent(chatId, a.botId, messageFromBot)
	}

	if err != nil {
		a.writeServerError(w, errors.New(errPrefix+err.Error()), "Automatic reply failed.")
		return
	}

	a.writeSuccess(w, "Automatic reply sent.")
}
