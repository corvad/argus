package main

import (
	"flag"

	"github.com/corvad/argus/util"
)

func main() {
	host := flag.String("host", "localhost:8080", "server host")
	flag.Parse()
	l, err := util.NewLoggerServer(*host)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		select {}
	}
}
