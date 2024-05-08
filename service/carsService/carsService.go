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

	carsIDStr := ctx.Param("id")
	carsID, err := strconv.Atoi(carsIDStr)
	if err != nil {
		log.Printf("[ERROR] Failed to convert cars ID to integer: %v", err)
		result = helpers.ResponseJSON(false, constans.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	cars, err := s.service.CarsRepo.FindCarsByID(carsID)
	if err != nil {
		log.Printf("[ERROR] Failed to find cars by ID %d: %v", carsID, err)
		result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	

	err = s.service.RedisRepo.InsertDataRedis(carsIDStr, cars)
	if err != nil {
		log.Printf("[ERROR] Failed to insert data into Redis for cars ID %d: %v", carsID, err)
		result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	log.Printf("[INFO] Successfully retrieved and cached cars data for ID %d", carsID)
	result = helpers.ResponseJSON(true, constans.SUCCESS_CODE, constans.EMPTY_CODE, cars)
	return ctx.JSON(http.StatusOK, result)
}
