package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func jmain() {
	resp, err := http.Get("https://www.youtube.com/results?search_query=golang+http+package")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	log.Printf(sb)
}
