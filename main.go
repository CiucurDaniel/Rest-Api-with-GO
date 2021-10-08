package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type Article struct{
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

// Articles global variable to hold a list of articles
var Articles []Article

// handle request to root URL "/"
func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Println("Endpoint reached: home page")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("Endpoint reached: /articles")
}

// function that will match the path in the URL with defined function handler
func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main(){

	// create some articles
	Articles = []Article{
		Article{Title: "Programming Java", Description: "A tutorial for Java"},
		Article{Title: "Computer Science", Description: "An introduction to Computer Science"},
	}

	// start server and handle incoming requests
	handleRequests()
}