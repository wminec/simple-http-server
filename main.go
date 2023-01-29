package main

import (
	"flag"
	"golang/simple-http-server/server"
	"log"
	//	"github.com/JakWai01/http-server/pkg/server"
)

func main() {
	port := flag.String("port", "8080", "Port to listen to")
	flag.Parse()

	listeningPort := ":" + *port
	log.Println(listeningPort)

	httpServer := server.NewHTTPServer(listeningPort)

	if err := httpServer.Open(); err != nil {
		log.Fatal("could not open httpServer", err)
	}
}
