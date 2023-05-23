package main

import (
	"context"
	"time"

	"github.com/go75/udpx/mod"
)

func main() {

	// regist write handler
	eng.Put(0, func(r mod.Request) {
		println(0)
		if !rd.SetNX(context.Background(), string(r.Payload), r.Addr.String(), time.Minute << 2).Val() {
			conn.WriteToUDP([]byte("err"), r.Addr)
			return
		}
		if !rd.SetNX(context.Background(), r.Addr.String(), string(r.Payload), time.Minute << 2).Val() {
			rd.Del(context.Background(), string(r.Payload))
			conn.WriteToUDP([]byte("err"), r.Addr)
			return
		}
		eng.Send([]byte("ok"), r.Addr)
	})
	
	// regist read handler
	eng.Put(1, func(r mod.Request) {
		println(1)
		res, err := rd.Get(context.Background(), string(r.Payload)).Result()
		if err != nil {
			conn.WriteToUDP([]byte{}, r.Addr)
			return
		}
		eng.Send([]byte(res), r.Addr)
	})

	// start server
	eng.Run()
}