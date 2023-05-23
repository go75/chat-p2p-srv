package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/go75/udpx/engine"

	"github.com/go-redis/redis/v8"
)

func init() {
	fileData, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileData, config)
	if err != nil {
		panic(err)
	}

	eng, err = engine.New(config.Server.Addr, 2, 2)
	if err != nil {
		panic(err)
	}
	rd = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: "",
		DB:       1,
	})

	pong, err := rd.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	} else {
		log.Println("redis inited, ", pong)
	}
}