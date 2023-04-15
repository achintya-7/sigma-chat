package sockets

import (
	"log"
	"net/http"
)

func Start() {
	server := Setup()

	log.Println("Starting socket server")

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving on localhost:4001")
	log.Fatal(http.ListenAndServe(":4001", nil))
}
