package repository

import "crud-redis/models"

type PingRepositoryInterface interface {
	Ping() (string, error)
}

type CarsRepositoryInterface interface {
	FindCarsByID(id int) (models.Cars, error)
}