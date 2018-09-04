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

	response_text := ""

	if request.RequestData.Command == "" {
		response_text = "Я буду повторять все ваши фразы, скажите что-нибудь"
	} else {
		response_text = request.RequestData.Command
	}

	response := alice.Response{
		ResponseData: alice.ResponseData{Text: response_text},
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
