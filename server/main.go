package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	eventsource "gopkg.in/antage/eventsource.v1"

	"github.com/gorilla/mux"
)

type config struct {
	static    string
	port      int
	events    chan event
	residents string
}

type event struct {
	Type string
	Data interface{}
}

type messageEvent struct {
	Message string `json:"message"`
}

func main() {
	log.SetPrefix("Homenum Revelio: ")
	c := parseArgs()
	c.events = make(chan event)

	fileServer := http.FileServer(http.Dir(c.static))

	r := mux.NewRouter()
	r.Handle("/", WhoIsHomeHandler(c.static))
	r.Handle("/updates", updatesHandler(c.events))
	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fileServer))

	beginScanning(c)

	log.Println("Listening on localhost:", c.port)
	port := fmt.Sprintf(":%d", c.port)
	log.Fatal(http.ListenAndServe(port, Log(r)))
}

// Logs requests made to the server
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// WhoIsHomeHandler displays a list of people on the network
func WhoIsHomeHandler(static string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, static+"/index.html")
	})
}

func updatesHandler(events chan event) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		es := eventsource.New(nil, nil)

		go func() {
			for event := range events {
				packet, err := json.Marshal(event.Data)
				if err != nil {
					log.Printf("JSON serialization failed, %v", err)
				} else {
					es.SendEventMessage(string(packet), event.Type, "")
				}
			}
		}()

		es.ServeHTTP(w, r)
	})
}

func beginScanning(c *config) {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for _ = range ticker.C {
			people := checkHouse()
			message := messageEvent{strings.Join(people, ", ")}
			c.events <- event{
				Type: "message",
				Data: message,
			}
		}
	}()
}

func checkHouse() []string {
	return []string{"Elliot", "Florian"}
}

func parseArgs() (c *config) {
	c = new(config)
	flag.IntVar(&c.port, "p", 8080, "Port for the server")
	flag.StringVar(&c.static, "static", "./client", "")
	flag.StringVar(&c.residents, "r", "./residents.yaml", "")
	flag.Parse()

	return
}
