package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	fmt.Println("Hello world!")
	server, _ := zmq.NewSocket(zmq.REP)
	server.Bind("tcp://0.0.0.0:6666")

	for {
		currentTime := time.Now()
		msg, err := server.RecvMessage(0)
		if err != nil {
			break
		}
		fmt.Printf("\n Message: %s, Current Time: %s", msg, currentTime)
		server.Send("OK", 0)
	}
}
