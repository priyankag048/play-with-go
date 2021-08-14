package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage")
	fmt.Println("Endpoint Hit:homePage")
}

func operateOnSingle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	key := params["id"]

	switch r.Method {
	case "GET":
		for _, article := range articles {
			if article.Id == key {
				json.NewEncoder(w).Encode(article)
			}
		}
	case "PUT":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var updatedArticle Article
		err := json.Unmarshal([]byte(reqBody), &updatedArticle)
		if err != nil {
			fmt.Fprintf(w, "Error occured %+v", err)
			return
		}
		articles = append(articles, updatedArticle)
		fmt.Fprintf(w, "Updated new article %+v", updatedArticle)
	case "DELETE":
		for index, article := range articles {
			if article.Id == key {
				articles = append(articles[:index], articles[index+1:]...)
			}
		}
	}

}

func operateOnArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		var newArticle Article
		err := json.Unmarshal([]byte(reqBody), &newArticle)
		if err != nil {
			fmt.Fprintf(w, "Error occured %+v", err)
		}
		articles = append(articles, newArticle)
		fmt.Fprintf(w, "%+v", string(newArticle.Id))
	case "GET":
		fmt.Println("Endpoint Hit: returnAllArticles")
		json.NewEncoder(w).Encode(articles)
	}

}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", operateOnArticle).Methods("POST", "GET")
	myRouter.HandleFunc("/articles/{id}", operateOnSingle).Methods("GET", "DELETE", "PUT")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
