package repository

import "crud-redis/models"

type PingRepositoryInterface interface {
	Ping() (string, error)
}

type CarsRepositoryInterface interface {
	GetCarsByIndex(id int) (models.Cars, error)
}