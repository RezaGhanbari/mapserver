package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"mapserver/cnst"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// Pool : new redis pool
	Pool *redis.Pool
)

// Init function for creating new redis pool
func Init(redisHost, redisPort string) *redis.Pool {
	redisURL := cnst.EMPTY
	if redisHost == cnst.EMPTY {
		redisURL = cnst.COLON + redisPort
	} else {
		redisURL = redisHost + cnst.COLON + redisPort
	}
	Pool = newPool(redisURL)
	cleanupHook()
	return Pool
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		err := Pool.Close()
		if err != nil {
			log.Println(err)
			return
		}
		os.Exit(0)
	}()
}

// Ping redis if it is alive or not
func Ping() error {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

// Get function in redis
func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	var data []byte
	data, err := redis.Bytes(conn.Do(cnst.GET, key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

// Set function in redis
func Set(key string, value []byte) error {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	_, err := conn.Do(cnst.SET, key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func exists(key string) (bool, error) {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

// Delete function in redis
func Delete(key string) error {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	_, err := conn.Do(cnst.DEL, key)
	return err
}

func getKeys(pattern string) ([]string, error) {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	iter := 0
	keys := []string{}

	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)
		if iter == 0 {
			break
		}
	}
	return keys, nil
}

func incr(counterKey string) (int, error) {
	conn := Pool.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	return redis.Int(conn.Do(cnst.INCR, counterKey))
}

