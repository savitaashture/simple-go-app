package main

import (
"fmt"
"log"
"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Println("My first app")
	fmt.Fprintf(w, "My first app with URL %s!", r.URL.Path[1:])
	time.Sleep(60 * time.Second)
}

func main() {
        fmt.Println("Its coming herre111111111111111111111111111111")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
