package carsService

import (
	"crud-redis/constans"
	"crud-redis/helpers"
	"crud-redis/models"
	"crud-redis/service"
	"encoding/json"
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

func (s carsService) GetAllCars(ctx echo.Context) error {
	var result models.Response

	// Check if cars data exists in Redis
	carsRedis, err := s.service.RedisRepo.GetDataRedis("cars")
	if err != nil {
		// Data not found in Redis, fetch from repository
		cars, err := s.service.CarsRepo.GetAllCars()
		if err != nil {
			log.Println("[ERROR] Failed to get all cars ", err)
			result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusInternalServerError, result)
		}
		// Insert data into Redis for caching
		err = s.service.RedisRepo.InsertDataRedis("cars", cars)
		if err != nil {
			log.Println("[ERROR] Failed to insert data into Redis for cars", err)
			result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		log.Println("[INFO] Successfully retrieved get all cars from repository")
		result = helpers.ResponseJSON(true, constans.SUCCESS_CODE, constans.EMPTY_CODE, cars)
		return ctx.JSON(http.StatusOK, result)
	}

	// Data found in Redis, use cached data
	var cars []models.Cars
	err = json.Unmarshal([]byte(carsRedis), &cars)
	if err != nil {
		log.Println("[ERROR] Failed to unmarshal cars data from Redis", err)
		result = helpers.ResponseJSON(false, constans.SYSTEM_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	log.Println("[INFO] Successfully retrieved get all cars from Redis")
	result = helpers.ResponseJSON(true, constans.SUCCESS_CODE, constans.EMPTY_CODE, cars)
	return ctx.JSON(http.StatusOK, result)
}

