package service

import (
	"crud-redis/repository"
	"database/sql"
)

type Service struct {
	DB       *sql.DB
	PingRepo repository.PingRepositoryInterface
	CarsRepo repository.CarsRepositoryInterface
}

func NewService(DB *sql.DB, pingRepo repository.PingRepositoryInterface, carsRepo repository.CarsRepositoryInterface) Service {
	return Service{
		DB:       DB,
		PingRepo: pingRepo,
		CarsRepo: carsRepo,
	}

}
