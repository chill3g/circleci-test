package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		q := r.URL.Query()
		fmt.Fprintf(w, "hello %s\n", q["name"][0])
	})
	return mux
}

func main() {
	s := &http.Server{
		Addr:    ":4000",
		Handler: Handler(),
	}
	log.Println("server is listening")
	log.Fatal(s.ListenAndServe())
}
