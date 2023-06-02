package main

import (
"fmt"
"log"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Println("My first app")
	fmt.Fprintf(w, "Welcome to KCD 2023")
	//time.Sleep(60 * time.Second)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
