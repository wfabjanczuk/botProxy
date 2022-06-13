package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Install(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("Installation request received:", body)

	err := createBot("Onboarding bot")

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not create the bot."))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Installation successful."))
	}
}

type createBotPayload struct {
	Name          string `json:"name"`
	OwnerClientId string `json:"owner_client_id"`
}

func createBot(name string) error {
	ownerClientId := os.Getenv("CLIENT_ID")
	if len(ownerClientId) == 0 {
		return errors.New("client id not set")
	}

	payload, err := json.Marshal(&createBotPayload{
		Name:          name,
		OwnerClientId: ownerClientId,
	})

	if err != nil {
		return err
	}

	request, err := getApiRequest(
		"POST",
		"https://api.labs.livechatinc.com/v3.4/configuration/action/create_bot",
		string(payload),
	)

	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("creating bot failed")
	}

	return nil
}
