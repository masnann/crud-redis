package routes

import (
	"crud-redis/service"
	"crud-redis/service/pingService"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, srv service.Service) {

	public := e.Group("/api/v1/public")
	ping := pingService.NewPingService(srv)
	pingGroup := public.Group("/ping")
	pingGroup.GET("/", ping.PingHandler)
}
