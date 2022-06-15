package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (a *app) Reply(w http.ResponseWriter, r *http.Request) {
	errPrefix := "replying to webhook request failed: "

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.writeClientError(w, fmt.Errorf("%s%w", errPrefix, err), "Invalid request body.")
		return
	}

	type postback struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	}
	type event struct {
		Postback postback `json:"postback"`
	}
	type payload struct {
		ChatId string `json:"chat_id"`
		Event  event  `json:"event"`
	}
	type incomingRequest struct {
		Payload   payload `json:"payload"`
		SecretKey string  `json:"secret_key"`
	}

	incoming := &incomingRequest{}
	err = json.Unmarshal(body, incoming)
	if err != nil {
		a.writeClientError(w, fmt.Errorf("%s%w", errPrefix, err), "Invalid JSON.")
		return
	}

	if incoming.SecretKey != a.conf.SecretKey {
		a.writeClientError(w, fmt.Errorf("%sinvalid secret key", errPrefix), "Invalid Secret.")
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
		a.writeServerError(w, fmt.Errorf("%s%w", errPrefix, err), "Automatic reply failed.")
		return
	}

	a.writeSuccess(w, "Automatic reply sent.")
}
