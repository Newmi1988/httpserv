package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Newmi1988/httpserv/server"
)

func main() {
	port := flag.String("p", "8080", "port")

	// get response as json string (Marshal)
	http.HandleFunc("/info", server.Info)
	// get response as stream of bytes
	http.HandleFunc("/post", server.PostStruct)

	serverPort := ":" + *port

	log.Println("** Service Started on Port " + *port + " **")

	http.ListenAndServe(serverPort, nil)
}
