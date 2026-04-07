package main

import (
	"flag"
	"time"

	"github.com/corvad/argus/util"
)

func main() {
	host := flag.String("host", "localhost:8080", "server host")
	service := flag.String("service", "service", "service name")
	flag.Parse()
	l, err := util.NewLogger(*host, *service)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for i := 0; i < 5; i++ {
		err = l.Log(util.INFO, "This is a test log message")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
