package app

import (
	"challenge-serasa/api/controller"
	"challenge-serasa/api/database"
	"challenge-serasa/api/handlers/integration"
	"challenge-serasa/api/handlers/negativations"

	httping "github.com/ednailson/httping-go"
	log "github.com/sirupsen/logrus"
)

type App struct {
	server    httping.IServer
	closeFunc httping.ServerCloseFunc
}

func LoadApp(cfg Config) (*App, error) {
	var app App
	db, err := database.NewDatabase(cfg.Database.Config)
	if err != nil {
		return nil, err
	}
	negativationColl, err := db.Collection(cfg.Database.NegativationCollection)
	if err != nil {
		return nil, err
	}
	controller := controller.NewController(negativationColl, cfg.MainframeUrl, cfg.Passphrase, cfg.Key)
	app.server = loadServer(controller)
	return &app, nil
}

func (a *App) Run() <-chan error {
	closeFunc, chErr := a.server.RunServer()
	a.closeFunc = closeFunc
	return chErr
}

func (a *App) Close() {
	err := a.closeFunc()
	if err != nil {
		log.WithField("error", err.Error()).Errorf("failed to close func")
	}
}

func loadServer(ctrl *controller.Controller) httping.IServer {
	server := httping.NewHttpServer("", 8082)
	integrationHandler := integration.NewHandler(*ctrl)
	server.NewRoute(nil, "/v1/integration").POST(integrationHandler.Handle)
	negativationsHandler := negativations.NewHandler(*ctrl)
	server.NewRoute(nil, "/v1/negativations/:customerDocument").GET(negativationsHandler.Handle)
	return server
}
