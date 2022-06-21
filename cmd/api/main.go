package main

import (
	"github.com/joho/godotenv"
	"github.com/wfabjanczuk/botProxy/internal/config"
	"github.com/wfabjanczuk/botProxy/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	app := handlers.NewApp(getConfig())

	http.HandleFunc("/install", app.Install)
	http.HandleFunc("/reply", app.Reply)

	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}

func getConfig() config.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	conf := config.Config{}

	conf.ClientId = os.Getenv("CLIENT_ID")
	if len(conf.ClientId) == 0 {
		log.Fatal("client id not set")
	}

	conf.SecretKey = os.Getenv("SECRET_KEY")
	if len(conf.SecretKey) == 0 {
		log.Fatal("secret key not set")
	}

	conf.BaseAccountsUrl = os.Getenv("BASE_ACCOUNTS_URL")
	if len(conf.BaseAccountsUrl) == 0 {
		log.Fatal("base accounts url not set")
	}

	conf.BaseApiUrl = os.Getenv("BASE_API_URL")
	if len(conf.BaseApiUrl) == 0 {
		log.Fatal("base api url not set")
	}

	conf.BaseAppUrl = os.Getenv("BASE_APP_URL")
	if len(conf.BaseAppUrl) == 0 {
		log.Fatal("base app url not set")
	}

	return conf
}
