// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"week04/api"
	"week04/internal/biz"
	"week04/internal/data"
	"week04/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	db, cleanup, err := data.NewPostgres()
	if err != nil {
		return nil, nil, err
	}
	client, cleanup2, err := data.NewREDIS()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dao, cleanup3, err := data.NewDao(db, client)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serverRepository := data.NewCloudServerRepo(dao)
	cloudServerRepo, cleanup4, err := biz.NewCloudServerRepo(serverRepository)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cloudService, cleanup5, err := service.NewCloudService(cloudServerRepo)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, cleanup6, err := NewHTTP(cloudService)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup7, err := NewApp(cloudService, server)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

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
