package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage endpoint hit!!!!!")
}

func allArticals(w http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{Title: "Testing", Desc: "Description", Content: "helllo"}}
	fmt.Fprint(w, "Article endpoint Hit endpoint hit!!!")
	json.NewEncoder(w).Encode(articles)
}

func handlerequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticals)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {

	handlerequest()
}
