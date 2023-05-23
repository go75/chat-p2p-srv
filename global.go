package main

import (
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/go75/udpx/engine"
)

var rd *redis.Client
var conn *net.UDPConn
var config = new(Config)
var eng *engine.Engine