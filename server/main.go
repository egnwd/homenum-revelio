package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
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
	People []person `json:"people"`
}

type person struct {
	Mac    string `json:"mac"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func main() {
	log.SetPrefix("Homenum Revelio: ")
	c := parseArgs()
	c.events = make(chan event)

	fileServer := http.FileServer(http.Dir(c.static))

	r := mux.NewRouter()
	r.Handle("/", WhoIsHomeHandler(c.static))
	r.Handle("/updates", updatesHandler(c))
	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fileServer))

	beginScanning(c)

	log.Println("Listening on localhost:", c.port)
	port := fmt.Sprintf(":%d", c.port)
	log.Fatal(http.ListenAndServe(port, Log(r)))
}

// Log requests made to the server
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

func updatesHandler(c *config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		es := eventsource.New(nil, nil)

		go func() {
			for event := range c.events {
				packet, err := json.Marshal(event.Data)
				if err != nil {
					log.Printf("JSON serialization failed, %v", err)
				} else {
					es.SendEventMessage(string(packet), event.Type, "")
				}
			}
		}()

		updateStatus(c)
		es.ServeHTTP(w, r)
	})
}

func beginScanning(c *config) {
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for _ = range ticker.C {
			updateStatus(c)
		}
	}()
}

func updateStatus(c *config) {
	c.events <- event{
		Type: "message",
		Data: checkHouse(c.residents),
	}
}

func checkHouse(residents string) messageEvent {
	cmd := exec.Command("./bin/homenum_revelio", residents)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var group messageEvent
	if err := json.NewDecoder(stdout).Decode(&group); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return group
}

func randomBool() bool {
	return int(rand.Float32()+0.5) == 1
}

func parseArgs() (c *config) {
	c = new(config)
	flag.IntVar(&c.port, "p", 8080, "Port for the server")
	flag.StringVar(&c.static, "static", "./client", "")
	flag.StringVar(&c.residents, "r", "./residents.yaml", "")
	flag.Parse()

	if _, err := os.Stat(c.residents); os.IsNotExist(err) {
		log.Fatal(err)
	}

	return
}
