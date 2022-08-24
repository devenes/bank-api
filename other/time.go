package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	// routers
	http.HandleFunc("/api/time", getCurrentTime)

	// server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	w.Header().Add("Content-Type", "application/json")

	// parse query params from the request
	params := r.URL.Query()
	tz, ok := params["tz"]
	if !ok || len(tz) == 0 {
		res = map[string]string{
			"UTC": time.Now().String(),
		}
		json.NewEncoder(w).Encode(res)
	}

	zones := strings.Split(tz[0], ",")
	for _, v := range zones {
		loc, err := time.LoadLocation(v)
		if err != nil {
			log.Printf("timezone %s not supported", v)
			w.WriteHeader(400)
			return
		}
		res[loc.String()] = time.Now().In(loc).String()
	}
	json.NewEncoder(w).Encode(res)
}
