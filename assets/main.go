package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("Displaying All Posts")
	json.NewEncoder(w).Encode(posts)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

//register functions
func handleRequest(){

	http.HandleFunc("/", homePage)
	http.HandleFunc("/posts", allPosts)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest();
}
