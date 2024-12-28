package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"strconv"
	"time"

	taskforce "github.com/utkarsh-singh1/project/go/restfuljsonapi/internal"
)

type taskServer struct {

	store *taskforce.TaskStore

}

func newServer() *taskServer {

	store := taskforce.New()

	taskserver := taskServer{store: store,}

	return &taskserver
	
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, req *http.Request ){

	type RequestTask struct {
		Text string `json:"text"`
		Due time.Time `json:"due"`
		Tag []string `json:"tag"`
	}

	type ResponseID struct {
		Id int `json:"id"`
	}


	// Request Section of code

	// This section make sure the request body contains only Json as the input fro the user side
	contentType := req.Header.Get("Content-type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediaType != "application/json" {
		http.Error(w, "expected Content-Type to be application/json", http.StatusUnsupportedMediaType)
	}

	
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var rt RequestTask
	if err := dec.Decode(&rt) ; err != nil {

		http.Error(w,err.Error() , http.StatusBadRequest)
		return
	}

	//Response Section of Code

	id := ts.store.CreateTask(rt.Text, rt.Tag, rt.Due)
	js ,err := json.Marshal(ResponseID{Id: id})
	if err != nil{

		http.Error(w, err.Error(), http.StatusInternalServerError)
		
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	
}

func (ts *taskServer) GetAllHandler (w http.ResponseWriter, req *http.Request) {

	allTasks := ts.store.GetAllTask()

	js , err := json.Marshal(allTasks)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func (ts *taskServer) GetTaskHandler (w http.ResponseWriter, req *http.Request) {

	id,err := strconv.Atoi(req.PathValue("id"))
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	task,err := ts.store.GetTask(id)
	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
		
	}

	js, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")

	w.Write(js)

}

func (ts *taskServer) GetTagTaskHandler (w http.ResponseWriter, req *http.Request) {

	tag := req.PathValue("tag")

	js, err := json.Marshal(tag)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func (ts *taskServer) GetDueTaskHandler(w http.ResponseWriter, req *http.Request) {

	badTaskRequest := func () {

		http.Error(w, fmt.Sprintf("Wanted the data in the format /due/<year>/<month>/<day> but got in the form of %v", req.URL.Path), http.StatusBadRequest)

	}

	year ,errYear := strconv.Atoi(req.PathValue("year"))
	month, errMonth := strconv.Atoi(req.PathValue("month"))
	date, errDate := strconv.Atoi(req.PathValue("date"))

	if errYear != nil || errMonth != nil || errDate != nil || month < int(time.January) || month > int(time.December) {

		badTaskRequest()
		return
	}

	task := ts.store.GetTaskByDueDate(year, time.Month(month), date)

	js, err := json.Marshal(task)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Tyep", "application/json")
	w.Write(js)
	
}


func (ts *taskServer) DeletetaskHandler (w http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = ts.store.DeleteTask(id)
	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (ts *taskServer) DeleteAllHandler (w http.ResponseWriter, req *http.Request) {

	ts.store.DeleteAllTask()

}


func main() {


	mux := http.NewServeMux()
	server := newServer()

	mux.HandleFunc("POST /task/", server.createTaskHandler)
	mux.HandleFunc("GET /task/", server.GetAllHandler)
	mux.HandleFunc("GET /task/{id}", server.GetTaskHandler)
	mux.HandleFunc("GET /task/{tag}", server.GetTagTaskHandler)
	mux.HandleFunc("GET /due/{year}/{month}/{day}", server.GetDueTaskHandler)
	mux.HandleFunc("DELETE /task/", server.DeleteAllHandler)
	mux.HandleFunc("DELETE /task/{id}", server.DeletetaskHandler)
	
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVEPORT"), mux))
	

}
