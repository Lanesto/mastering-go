package main

import (
	"fmt"
	"net/http"
	"os"
)

func checkStatusOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Fine!`)
}

func statusNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.statusNotFound)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := ":8080"
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Using default port number:", PORT)
	} else {
		PORT = ":" + args[1]
	}

	http.HandleFunc("/checkStatusOK", checkStatusOK)
	http.HandleFunc("/statusNotFound", statusNotFound)
	http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
