package main

import (
	"fmt"
	"net/http"
	"flag"
	"github.com/Newmi1988/httpserv/server"
)

func main() {
	fmt.Println("-- go http server --")
	port := flag.String("p","8080","port")

	// get response as json string (Marshal)
	http.HandleFunc("/info",server.Info)
	// get response as stream of bytes
	http.HandleFunc("/post",server.PostStruct)

	serverPort := ":" + *port

	http.ListenAndServe(serverPort, nil)
}
