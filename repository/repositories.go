package repository

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	RDB *redis.Client
	DB  *sql.DB
}

func NewRepository(rdb *redis.Client, db *sql.DB) Repository {
	return Repository{
		RDB: rdb,
		DB: db,
	}
}
