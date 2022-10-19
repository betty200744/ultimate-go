package main

import (
	"fmt"
	"log"
	"time"

	"github.com/flyaways/pool"
	"google.golang.org/grpc"
)

func main() {
	options := &pool.Options{
		InitTargets:  []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"},
		InitCap:      5,
		MaxCap:       30,
		DialTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	p, err := pool.NewGRPCPool(options, grpc.WithInsecure()) //for grpc
	//p, err := pool.NewRPCPool(options) 			//for rpc
	//p, err := pool.NewTCPPool(options)			//for tcp

	if err != nil {
		log.Printf("%#v\n", err)
		return
	}

	if p == nil {
		log.Printf("p= %#v\n", p)
		return
	}

	defer p.Close()

	//todo
	//danamic update targets
	//options.Input()<-&[]string{}
	var conn *grpc.ClientConn
	for i := 0; i < 3; i++ {
		conn, err = p.Get()
		if err != nil {
			log.Printf("%#v\n", err)
			return
		}
		fmt.Println(conn.Target())
		p.Put(conn)
	}

	conn, err = p.Get()
	if err != nil {
		log.Printf("%#v\n", err)
		return
	}
	fmt.Println(conn.Target())
	defer p.Put(conn)
	//todo
	//conn.DoSomething()

	log.Printf("len=%d\n", p.IdleCount())
}
