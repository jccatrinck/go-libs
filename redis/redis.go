package redis

import (
	"log"

	"github.com/jccatrinck/go-libs/env"

	"github.com/garyburd/redigo/redis"
)

// Pool maintains a pool of connections
var Pool redis.Pool

// Load redis
func Load() (err error) {
	host := env.Get("REDIS_HOST", "localhost")
	port := env.Get("REDIS_PORT", "6379")

	Pool = redis.Pool{
		MaxIdle:   50,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				log.Panic(err)
			}
			return c, err
		},
	}

	return
}

// Exec redis query
func Exec(task func(redis.Conn) (interface{}, error)) (interface{}, error) {
	conn := Pool.Get()
	defer conn.Close()
	return task(conn)
}

func GetString(key string) (string, error) {
	v, err := Exec(func(conn redis.Conn) (interface{}, error) {
		return redis.String(conn.Do("GET", key))
	})

	s, _ := v.(string)

	return s, err
}

func SetString(key string, value string) error {
	_, err := Exec(func(conn redis.Conn) (interface{}, error) {
		return conn.Do("SET", key, value)
	})

	return err
}
