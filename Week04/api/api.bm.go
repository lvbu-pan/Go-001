package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

var PathDemoCreateCloudServer = "/demo/create_cloud_server"

// DemoBMServer is the server API for Demo service.
type DemoBMServer interface {
	CreateCloudServer(ctx context.Context, req *CloudServerReq) (resp *CloudServerResp)
}

var DemoSvc DemoBMServer

func demoCreateCloudServer(g *gin.Context) {
	p := new(CloudServerReq)
	if err := g.ShouldBindJSON(p); err != nil {
		g.AbortWithStatus(400)
		return
	}
	a := DemoSvc
	fmt.Println(a)
	resp := DemoSvc.CreateCloudServer(context.Background(), p)
	g.JSON(200, resp)
}

func RegisterDemoBMServer(e *gin.Engine, server DemoBMServer) {
	DemoSvc = server
	e.POST("/demo/create_cloud_server", demoCreateCloudServer)
}
