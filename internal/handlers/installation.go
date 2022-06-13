package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Install(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("Installation request received:", body)
}
