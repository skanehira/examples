package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	pid := os.Getpid()
	log.Printf("server pid: %d\n", pid)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	log.Fatal(http.ListenAndServe(":9999", nil))
}
