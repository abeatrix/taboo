package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Post struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Posts []Post

func allPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{Title: "A Test Tiltle", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("All Posts Are Here")
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Post added Successfully")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage Endpoint Landed")
}

//register functions
func handleRequest(){

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/posts", allPosts).Methods("GET")
	myRouter.HandleFunc("/posts", addPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest();
}
