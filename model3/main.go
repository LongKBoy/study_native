package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	fmt.Println("starting http server ....")
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	defer printEvent(w, r)
	addHeader(w, r)
	w.WriteHeader(200)
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	defer printEvent(w, r)
	addHeader(w, r)
	w.WriteHeader(200)
	io.WriteString(w, "root ok\n")
}
func addHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, fmt.Sprintf("%v", v))
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
}
func printEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("client IP: %s,response code %d\n", strings.Split(r.RemoteAddr, ":")[0], 200)
	for k, v := range w.Header() {
		fmt.Printf("%v=%v\n", k, v)
	}
}
