package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("usage\n pingclient http://localhost:8080")
		log.Fatal("must provide addr to ping")
	}
	addr := os.Args[1]
	res, err := http.Get(addr)
	if err != nil {
		log.Fatalf("ping %s failed %v", addr, err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("can't read body %v", err)
	}
	log.Println(string(b))
}
