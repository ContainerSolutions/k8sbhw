package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/riccardomc/k8sbhw/webapp/datastore"
)

func getStoreHandler(dataStore datastore.Datastore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		switch r.Method {

		case "GET":
			w.Header().Add("Content-type", "application/json")
			err := json.NewEncoder(w).Encode(dataStore.Get())
			if err != nil {
				status = http.StatusInternalServerError
				w.WriteHeader(status)
			}

		case "PUT":
			record := datastore.Record{}
			err := json.NewDecoder(r.Body).Decode(&record)
			if err != nil {
				status = http.StatusBadRequest
				w.WriteHeader(status)
				fmt.Fprintln(w, err)
			} else {
				dataStore.Add(record)
				status = http.StatusCreated
				w.WriteHeader(status)
				fmt.Fprintln(w, "OK")
			}

		case "DELETE":
			record := datastore.Record{}
			err := json.NewDecoder(r.Body).Decode(&record)
			if err != nil {
				status = http.StatusBadRequest
				w.WriteHeader(status)
				fmt.Fprintln(w, err)
			} else {
				dataStore.Rem(record)
				status = http.StatusOK
				w.WriteHeader(status)
				fmt.Fprintln(w, "OK")
			}

		default:
			status = http.StatusMethodNotAllowed
			w.WriteHeader(status)
		}

		log(fmt.Sprintf("%s - %s %d", r.Method, r.URL.Path, status))
	}
}

func log(message string) {
	now := time.Now().Format("2006/01/02 15:04:05")
	hostname := os.Getenv("HOSTNAME")
	fmt.Printf("%s [%s] %s\n", now, hostname, message)
}

func main() {
	listen := ":9009"
	dataStore := datastore.NewSliceDataStore()
	dataStore.Init(0)
	http.HandleFunc("/store", getStoreHandler(dataStore))
	log(fmt.Sprintf("Listening on %s", listen))
	http.ListenAndServe(listen, nil)
}
