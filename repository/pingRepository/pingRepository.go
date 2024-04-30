package pingRepository

import (
	"crud-redis/repository"
	"fmt"
)

type PingRepository struct {
	repo repository.Repository
}

func NewPingRespository(repo repository.Repository) PingRepository {
	return PingRepository{
		repo: repo,
	}
}

func (r PingRepository) Ping() (string, error) {
	msg := "Hello, world!"
	fmt.Println(msg)
	return msg, nil
}
