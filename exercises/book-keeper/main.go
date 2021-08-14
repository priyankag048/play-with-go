package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model // add uuid

	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

var db *gorm.DB
var err error

func main() {
	// Loading environment variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Open the connection
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	// Close connection to database when main function finishes
	defer db.Close()

	// Make database migrations
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
	handleRequests()
}

// API routes
func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", healthcheck).Methods("GET")
	router.HandleFunc("/person", operateOnPersons).Methods("GET", "POST")
	router.HandleFunc("/person/{id}", operateOnPerson).Methods("GET", "DELETE")
	httpErr := http.ListenAndServe(":8001", router)
	if httpErr != nil {
		panic(err)
	}
	println("Running code after ListenAndServe")
}

// API controllers
func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthcheck success")
}
func operateOnPersons(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var persons []Person
		db.Find(&persons)
		json.NewEncoder(w).Encode(&persons)
		return
	case "POST":
		var person Person
		json.NewDecoder(r.Body).Decode(&person)
		createdPerson := db.Create(&person)
		err = createdPerson.Error
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(&person)
		}
		return
	}

}

func operateOnPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	var books []Book
	params := mux.Vars(r)
	id := params["id"]
	switch r.Method {
	case "GET":
		db.First(&person, id)
		db.Model(&person).Related(&books)
		json.NewEncoder(w).Encode(&person)
		return
	case "DELETE":
		db.First(&person, id)
		db.Delete(&person)
		json.NewEncoder(w).Encode(&person.ID)
		return
	}

}
