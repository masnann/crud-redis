package app

import (
	"crud-redis/repository"
	"crud-redis/repository/carsRepository"
	"crud-redis/repository/pingRepository"
	"crud-redis/repository/redisRepository"
	"crud-redis/service"
	"database/sql"
)

func SetupApp(DB *sql.DB, repo repository.Repository) service.Service {
	pingRepo := pingRepository.NewPingRespository(repo)
	carsRepo := carsRepository.NewCarsRepository(repo)
	redisRepo := redisRepository.NewRedisRepository(repo)

	service := service.NewService(DB, pingRepo, carsRepo, redisRepo)

	return service
}
