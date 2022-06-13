package handlers

import (
	"github.com/wfabjanczuk/botProxy/internal/config"
	"github.com/wfabjanczuk/botProxy/internal/requests"
	"log"
	"net/http"
)

type app struct {
	conf              config.Config
	client            *requests.Client
	installationSteps []func(w http.ResponseWriter) bool
	botId             string
}

func NewApp(conf config.Config) *app {
	a := &app{
		conf:   conf,
		client: requests.NewClient(conf),
	}
	a.installationSteps = []func(w http.ResponseWriter) bool{
		a.createBot, a.setRoutingStatus, a.registerWebhook, a.enableLicenseWebhooks,
	}

	return a
}

func (a *app) writeServerError(w http.ResponseWriter, err error, safeMessage string) bool {
	log.Println(err)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(safeMessage))

	return false
}

func (a *app) writeSuccess(w http.ResponseWriter, safeMessage string) bool {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(safeMessage))

	return true
}
