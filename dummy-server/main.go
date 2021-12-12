package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("dummy-server serving on http://localhost:5000/hello/sayHello")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
