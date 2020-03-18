package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "serving"
	}
	fmt.Fprintf(w, "Hello tekton %s!\n", target)
}

func main() {
	log.Print("Hello world started.")

	http.HandleFunc("/", envHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
