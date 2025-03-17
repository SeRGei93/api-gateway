package main

import (
	"log"
	"net/http"

	"auth/internal/app/store/mysql"
	dbv4 "github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

func main() {
	dbCfg := dbv4.NewConfig("root:123456@tcp(127.0.0.1:3306)/myapp_test?charset=utf8")

	manager := manage.NewDefaultManager()
	// use mysql token store
	store := dbv4.NewDefaultStore(dbCfg)
	defer store.Close()
	manager.MapTokenStorage(store)

	manager.MapClientStorage(mysql.NewMysqlClientStore(dbCfg))

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
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
