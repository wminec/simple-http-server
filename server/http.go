package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

// HTTPServer creates a http server and can be reached through the provided port
type HTTPServer struct {
	port string
}

// NewHTTPServer initializes variables
func NewHTTPServer(port string) *HTTPServer {
	return &HTTPServer{port}
}

// Open creates the http server
func (s HTTPServer) Open() error {
	http.HandleFunc("/", home)
	http.HandleFunc("/dnscheck", dnscheck)
	http.ListenAndServe(s.port, nil)

	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func dnscheck(w http.ResponseWriter, r *http.Request) {
	host := r.FormValue("host")
	resolver := net.Resolver{}
	ips, err := resolver.LookupIPAddr(r.Context(), host)
	if err != nil {
		fmt.Fprintf(w, "Failed to resolve %s: %v\n", host, err)
		return
	} else {
		fmt.Fprintf(w, "Resolved %s to %v\n", host, ips)
	}
}
