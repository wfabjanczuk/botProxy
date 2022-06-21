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
	installationSteps []func(w http.ResponseWriter, r *http.Request) bool
	botId             string
}

func NewApp(conf config.Config) *app {
	a := &app{
		conf:   conf,
		client: requests.NewClient(conf),
	}

	a.installationSteps = []func(w http.ResponseWriter, r *http.Request) bool{
		a.authorize,
		a.createBot,
		a.setRoutingStatus,
	}

	return a
}

func (a *app) writeServerError(w http.ResponseWriter, err error, safeMessage string) bool {
	return a.writeError(w, err, safeMessage, http.StatusInternalServerError)
}

func (a *app) writeClientError(w http.ResponseWriter, err error, safeMessage string) bool {
	return a.writeError(w, err, safeMessage, http.StatusBadRequest)
}

func (a *app) writeError(w http.ResponseWriter, err error, safeMessage string, statusCode int) bool {
	log.Println(err)

	w.WriteHeader(statusCode)
	w.Write([]byte(safeMessage))

	return false
}

func (a *app) writeSuccess(w http.ResponseWriter, safeMessage string) bool {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(safeMessage))

	return true
}
