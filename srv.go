package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, welcomeMessage(r.URL.Path[1:]))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "version 5")
	})
	log.Println("start serving on port 8080")
	http.ListenAndServe(":8080", nil)
}

func welcomeMessage(name string) string {
	return fmt.Sprintf("Hi %s, welcome to ECS!", name)
}
