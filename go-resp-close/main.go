package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func countPD(pid int) {
	de, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("fd count:", len(de))
}

func main() {
	pid := os.Getpid()
	log.Printf("client pid: %d\n", pid)
	for i := 0; i < 10; i++ {
		req, _ := http.NewRequest("POST", "http://localhost:9999", nil)
		// respのbodyを捨てる
		_, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		countPD(pid)
	}
}
