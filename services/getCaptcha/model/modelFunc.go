package model

import (
	"github.com/gomodule/redigo/redis"
)

func SaveImgCode(code, uuid string) error {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("setex", uuid, 60*5, code)
	if err != nil {
		return err
	}
	return nil
}
