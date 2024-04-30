package service

import (
	"crud-redis/repository"
	"database/sql"
)

type Service struct {
	DB       *sql.DB
	PingRepo repository.PingRepositoryInterface
}

func NewService(DB *sql.DB, pingRepo repository.PingRepositoryInterface) Service {
	return Service{
		DB:       DB,
		PingRepo: pingRepo,
	}

}
