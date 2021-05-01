package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homepage endpoint hit!!!!!")
}

func allArticals(w http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{Title: "Testing", Desc: "Description", Content: "helllo"}}
	fmt.Fprint(w, "Article endpoint Hit endpoint hit!!!")
	json.NewEncoder(w).Encode(articles)
}
func allArticals1(w http.ResponseWriter, r *http.Request) {
	articles := Articles{Article{Title: "Testing", Desc: "Description", Content: "helllo"}}
	fmt.Fprint(w, "Article endpoint Hit endpoint hit!!!----------------")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", allArticals)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func handleRequests1() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", allArticals1)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func main() {
	//handling multiple servers in go
	go handleRequests()
	handleRequests1()
}
