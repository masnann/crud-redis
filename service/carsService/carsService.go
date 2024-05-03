package carsService

import (
	"crud-redis/constans"
	"crud-redis/helpers"
	"crud-redis/models"
	"crud-redis/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type carsService struct {
	repo repository.CarsRepositoryInterface
}

func NewCarsService(repo repository.CarsRepositoryInterface) carsService {
	return carsService{
		repo: repo,
	}
}

func (s carsService) FindCarsByID(ctx echo.Context) error {
	var result models.Response

	carsID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Error convert id :", err)
		result = helpers.ResponseJSON(constans.VALIDATE_ERROR_CODE, err.Error(), false, nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	
}
