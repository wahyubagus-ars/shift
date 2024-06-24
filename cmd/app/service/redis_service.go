package service

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	redisService     *RedisServiceImpl
	redisServiceOnce sync.Once
)

type RedisService interface {
	GetCache(key string) (interface{}, error)
	PutCache(key string, data interface{}, c *gin.Context) error
	PutCacheBatch(data map[string]interface{}, c *gin.Context) error
}

type RedisServiceImpl struct {
	redisdb *redis.Client
}

func (r *RedisServiceImpl) GetCache(key string) (interface{}, error) {
	return "test", nil
}

func (r *RedisServiceImpl) PutCache(key string, data interface{}, c *gin.Context) error {
	err := r.redisdb.Set(c, key, data, 0).Err()
	if err != nil {
		log.Error("Error when try to put data to redis cache: ", err)
		return err
	}

	return nil
}

func (r *RedisServiceImpl) PutCacheBatch(data map[string]interface{}, c *gin.Context) error {
	err := r.redisdb.MSet(c, data).Err()
	if err != nil {
		log.Error("Error when try to put batch data to redis cache: ", err)
		return err
	}

	return nil
}

func ProvideRedisService(redis *redis.Client) *RedisServiceImpl {
	redisServiceOnce.Do(func() {
		redisService = &RedisServiceImpl{
			redisdb: redis,
		}
	})

	return redisService
}
