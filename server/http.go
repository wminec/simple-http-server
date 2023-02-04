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
	http.HandleFunc("/dnscheck", dnsCheck)
	http.HandleFunc("/rdnscheck", revDnsCheck)
	http.HandleFunc("/allreqheader", printAllReqHeader)

	http.ListenAndServe(s.port, nil)

	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func dnsCheck(w http.ResponseWriter, r *http.Request) {
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

func revDnsCheck(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("ip")
	resolver := net.Resolver{}
	host, err := resolver.LookupAddr(r.Context(), ip)
	if err != nil {
		fmt.Fprintf(w, "Failed to resolve %s: %v\n", ip, err)
		return
	} else {
		fmt.Fprintf(w, "resolved %s to %v\n", ip, host)
	}
}

func printAllReqHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s : %s\n", k, v)
	}
}
