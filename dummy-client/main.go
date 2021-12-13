package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/server/ask", askServer)
	log.Printf("dummy-client serving on http://localhost:5001/server/ask")
	log.Fatal(http.ListenAndServe(":5001", nil))

}

func askServer(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}

	resp, err := client.Get("http://localhost:5000/hello/sayHello")

	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}

	fmt.Printf("Got response %d: %s %s", resp.StatusCode, resp.Proto, string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server said " + string(body)))
}
