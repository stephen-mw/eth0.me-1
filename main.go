package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	// @readonly
	runHttpMsg     = "Runing eth0.me server on port [%s]: %d"
	rootHandlerMsg = "%s: %s %s request from %s.\n"
)

type Eth0Me struct{}

// Root handler
func (s *Eth0Me) rootHandler(res http.ResponseWriter, req *http.Request) {
	remote_addr := strings.Split(req.RemoteAddr, ":")[0]
	log.Printf(rootHandlerMsg, req.Host, req.Method, req.URL, remote_addr)
	fmt.Fprintf(res, remote_addr)
}

// Null handler do nothing but drop connection.
func (s *Eth0Me) nullHandler(res http.ResponseWriter, req *http.Request) {}

func (s *Eth0Me) Run(protocol string, port int) {
	r := mux.NewRouter()
	r.HandleFunc(`/`, s.rootHandler)
	r.HandleFunc(`/favicon.ico`, s.nullHandler)

	// Indicate port listening
	log.Printf(runHttpMsg, protocol, port)

	sockaddr := fmt.Sprintf(":%d", port)

	// Listen HTTP or HTTPS port
	switch protocol {
	case "HTTPS":
		log.Fatal(http.ListenAndServeTLS(sockaddr, "server.pem", "server.key", r))
	default:
		log.Fatal(http.ListenAndServe(sockaddr, r))
	}
}

func main() {
	ports := map[string]int{
		"HTTP":  8080,
		"HTTPS": 8081,
	}

	wg := &sync.WaitGroup{}

	// Start listing HTTP and HTTPS ports
	for protocol, port := range ports {
		wg.Add(1)
		go func(protocol string, port int) {
			new(Eth0Me).Run(protocol, port)
			wg.Done()
		}(protocol, port)
	}

	wg.Wait()
}
