package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


type Route struct {

	Name , Method, Pattern string
	HandlerFunc http.HandlerFunc 
}

type Routes []Route

var routes = Routes{

	Route{
		"Index", "GET", "/", Index,
	},

	Route{
		"TodoIndex", "GET", "/todos", TodoIndex,
	},
	
	Route{
		"TodoShow", "GET", "/todos/{todoId}", TodoShow,
	},
}

func NewRouter() *mux.Router{

	router := mux.NewRouter().StrictSlash(true)

	for _ , route := range routes {

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

type Todo struct {

	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type Todos []Todo

func main() {

	router := NewRouter()
		
	log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Welcome !")
}

func TodoIndex(w http.ResponseWriter, req *http.Request) {

	todos := Todos{
		Todo{Name: "Task1"},
		Todo{Name: "Task2"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	todoId := vars["todoId"]
	fmt.Fprintln(w , todoId)
}
