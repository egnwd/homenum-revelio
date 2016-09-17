package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type config struct {
	static string
	port   int
}

func main() {
	log.SetPrefix("Homenum Revelio: ")
	c := parseArgs()

	r := mux.NewRouter()
	r.Handle("/", WhoIsHomeHandler(c.static))

	log.Println("Listening on localhost:", c.port)
	port := fmt.Sprintf(":%d", c.port)
	log.Fatal(http.ListenAndServe(port, r))
}

// WhoIsHomeHandler displays a list of people on the network
func WhoIsHomeHandler(static string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, static+"/index.html")
	})
}

func parseArgs() (c *config) {
	c = new(config)
	flag.IntVar(&c.port, "p", 8080, "Port for the server")
	flag.StringVar(&c.static, "static", "./client", "")
	flag.Parse()

	return
}
