package carsService

import (
	"crud-redis/constans"
	"crud-redis/helpers"
	"crud-redis/models"
	"crud-redis/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type carsService struct {
	service service.Service
}

func NewCarsService(service service.Service) carsService {
	return carsService{
		service: service,
	}
}

func (s carsService) FindCarsByID(ctx echo.Context) error {
	var result models.Response

	carsID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("[SERVICE] Error convert id: ", err)
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	cars, err := s.service.CarsRepo.FindCarsByID(carsID)
	if err != nil {
		log.Println("[SERVICE] Error FindCarsByID: ", err)
		result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result = helpers.ResponseJSON(true, constans.SUCCESS_CODE, constans.EMPTY_CODE, cars)
	return ctx.JSON(http.StatusOK, result)
}
