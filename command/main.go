package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    	fmt.Fprintf(w, "Hello from Docker")
    })

    fmt.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}