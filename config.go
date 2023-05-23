package main

type Config struct {
	Server Server
	Redis Redis
}

type Server struct {
	Addr string `json:"Addr"`
}

type Redis struct {
	Addr string `json:"Addr"`
}