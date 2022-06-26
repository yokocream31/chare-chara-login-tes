package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    port := "8080"
    fmt.Printf("Server Listening on port %s\n", port)

    http.HandleFunc("/", handler)

    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("fatal err: ", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}