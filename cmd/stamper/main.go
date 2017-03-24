package main

import (
	"flag"
	"log"
	"os"

	"github.com/tombell/stamper/web"
)

func main() {
	host := flag.String("host", "127.0.0.1", "Host to bind to")
	port := flag.String("port", "8080", "Port to listen on")

	flag.Parse()

	l := log.New(os.Stderr, "[stamper] ", log.Lshortfile)
	l.Println("Starting...")

	var err error

	err = web.Run(*host, *port, l)
	if err != nil {
		l.Fatal(err)
	}
}
