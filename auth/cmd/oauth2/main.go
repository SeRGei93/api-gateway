package main

import (
	"auth/internal/app/store/mysql"
	"auth/internal/config"
	"fmt"
	dbv4 "github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
		return
	}

	dbCfg := dbv4.NewConfig(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name))

	manager := manage.NewDefaultManager()
	// use mysql token store
	store := dbv4.NewDefaultStore(dbCfg)
	defer store.Close()
	manager.MapTokenStorage(store)

	clientStore, err := mysql.NewMysqlClientStore(dbCfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return "000000", nil
	}

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
