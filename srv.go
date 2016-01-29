package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, welcomeMessage(r.URL.Path[1:]))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func welcomeMessage(name string) string {
	return fmt.Sprintf("Hi %s, welcome to ECS!", name)
}
