package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/r00takaspin/pingpong/alice"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var request alice.Request

	if r.Body == nil {
		http.Error(w, "Empty body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	response := alice.Response{
		ResponseData: alice.ResponseData{Text: request.RequestData.Command},
		Session: request.Session,
		Version: request.Version,
		EndSession: true,
	}

	res1B, _ := json.Marshal(response)

	w.Write([]byte(res1B))
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + string(port), nil))
}
