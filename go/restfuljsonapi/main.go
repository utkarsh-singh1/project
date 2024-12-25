package main

import (
	"encoding/json"
	"mime"
	"net/http"
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

func main() {

	

}
