package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return
	}
	defer conn.Close()

	reply, err := conn.Do("set", "key1", "value1")
	if err != nil {
		return
	}
	r, e := redis.String(reply, err)
	if e != nil {
		return
	}
	fmt.Println(r, e)

}
