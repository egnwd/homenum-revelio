package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	static := "./client"
	r.Handle("/", WhoIsHomeHandler(static))

	log.Fatal(http.ListenAndServe(":8080", r))
}

// WhoIsHomeHandler displays a list of people on the network
func WhoIsHomeHandler(static string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, static+"/index.html")
	})
}
