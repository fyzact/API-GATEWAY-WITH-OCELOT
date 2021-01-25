package main

import (
	"log"
	"net/http"
)

func main() {
	foo := func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello I am Foo service from GO "))
	}

	http.HandleFunc("/api/foo", foo)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
