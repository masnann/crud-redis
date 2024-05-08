package redisRepository

import (
	"context"
	"crud-redis/repository"
	"encoding/json"
)

type RedisRepository struct {
	repo repository.Repository
}

func NewRedisRepository(repo repository.Repository) RedisRepository {
	return RedisRepository{
		repo: repo,
	}
}

func (r RedisRepository) InsertDataRedis(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

   err = r.repo.RDB.Set(context.Background(), key, jsonData, 0).Err()
   if err != nil {
	 return err
   }

   return nil
}




