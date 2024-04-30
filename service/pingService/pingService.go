package pingService

import (
	"crud-redis/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingService struct {
	svc service.Service
}

func NewPingService(svc service.Service) PingService {
	return PingService{
		svc: svc,
	}
}

func (s PingService) PingHandler(c echo.Context) error {
    msg, err := s.svc.PingRepo.Ping()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "message": "Failed to ping data source",
            "error":   err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": msg,
    })
}