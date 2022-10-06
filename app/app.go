package app

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/robotxt/goboiler/app/rest"
	"github.com/sirupsen/logrus"
)

type App struct {
	Ctx    context.Context
	Router *mux.Router
	Log    *logrus.Logger
}

func (a *App) ActivateAPI() {
	a.Log.Info("Starting Rest API endpoints")
	rest_srv := rest.RestService{Ctx: a.Ctx, Router: a.Router, Log: a.Log}
	rest_srv.RegisterEndpoints()
}

func InitializeApp(ctx context.Context, router *mux.Router, log *logrus.Logger) *App {
	return &App{
		Ctx:    ctx,
		Router: router,
		Log:    log,
	}
}
