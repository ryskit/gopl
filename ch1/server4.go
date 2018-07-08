package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler4)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	v := r.URL.Query()
	if v == nil {
		return
	}
	for key, vs := range v {
		fmt.Fprintf(w, "%s = %s\n", key, vs[0])
	}
}
