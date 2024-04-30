package app

import (
	"crud-redis/repository"
	"crud-redis/repository/pingRepository"
	"crud-redis/service"
	"database/sql"
)

func SetupApp(DB *sql.DB, repo repository.Repository) service.Service {
	pingRepo := pingRepository.NewPingRespository(repo)

	service := service.NewService(DB, pingRepo)

	return service
}
