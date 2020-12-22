// +build wireinject

package di

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"log"
	"net/http"
	"time"
	"week04/api"
	"week04/internal/biz"
	"week04/internal/data"
	"week04/internal/service"
)

type App struct {
	srv *service.CloudService
	Hs  *http.Server
}

func NewApp(cloudService *service.CloudService, hs *http.Server) (*App, func(), error) {
	app := &App{srv: cloudService, Hs: hs}
	clean := func() {}
	return app, clean, nil
}

func NewHTTP(s api.DemoBMServer) (*http.Server, func(), error) {
	g := gin.Default()
	api.RegisterDemoBMServer(g, s)
	hSrv := &http.Server{
		Addr:    "127.0.0.1:9999",
		Handler: g,
	}
	clean := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		if err := hSrv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}
	return hSrv, clean, nil
}

func InitApp() (*App, func(), error) {
	wire.Build(data.Provider, biz.NewCloudServerRepo, service.Provider, NewHTTP, NewApp)
	return nil, nil, nil
}
