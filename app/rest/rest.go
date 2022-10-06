package rest

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/robotxt/goboiler/app/core"
	"github.com/robotxt/goboiler/app/repo"
	"github.com/sirupsen/logrus"
)

type RestService struct {
	Ctx     context.Context
	Router  *mux.Router
	Log     *logrus.Logger
	DB      *gorm.DB
	CoreSrv core.Service
}

func (srv RestService) RegisterEndpoints() {
	db_repo := repo.NewRepoService(srv.Log)
	db, err := db_repo.Connect()

	if err != nil {
		srv.Log.Fatal(err)
	}

	core_srv := core.NewCoreService(srv.Ctx, db, srv.Log)
	srv.CoreSrv = core_srv

	srv.Router.HandleFunc("/", srv.baseAPI).Methods(http.MethodGet)
	srv.Router.HandleFunc("/registration", srv.registrationAPI).Methods(http.MethodPost)
	srv.Router.HandleFunc("/login", srv.LoginAPI).Methods(http.MethodPost)

	// secure API endpoints
	secure_api := srv.Router.PathPrefix("/api/v1").Subrouter()
	secure_api.Use(srv.SecureRequest)
	secure_api.HandleFunc("/", srv.homeAPI).Methods(http.MethodGet)

}

func (srv RestService) baseAPI(w http.ResponseWriter, r *http.Request) {
	renderResponse(w, "Hello World. Welcome to the API", http.StatusOK)
}

func (srv RestService) homeAPI(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserCtxKey("AuthUser")).(AuthUser)
	renderResponse(w, "Hey, There! "+user.Email, http.StatusOK)
}
