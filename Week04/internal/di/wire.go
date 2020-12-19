//+build wireinject

package di

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"log"
	"net/http"
	"time"
	"week04/internal/data"
	"week04/api"
)

type App struct {
	HS *http.Server
	Db *sql.DB
	Rd *redis.Client
}

var Routes = map[string]func(ctx *gin.Context){
	"/create_asset": api.Hosts,
}

func NewHttp() (*http.Server, func(), error) {
	routes := gin.Default()

	registerRoutes(routes)

	hSrv := &http.Server{
		Addr:    "127.0.0.1:9999",
		Handler: routes,
	}
	clean := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		if err := hSrv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
		cancel()
	}
	return hSrv, clean, nil
}

// 注册路由函数
func registerRoutes(engine *gin.Engine) {
	for route, fn := range Routes {
		engine.POST(route, fn)
	}
}

func NewApp(d *sql.DB, r *redis.Client, h *http.Server) (*App, func(), error) {
	return &App{h, d, r}, func() {}, nil
}

func InitResource() (*App, func(), error) {
	wire.Build(data.Provider, NewHttp, NewApp)
	return nil, nil, nil
}
