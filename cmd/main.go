package main

import (
	"os"
	"runtime"
	"net/http"
	"encoding/json"
	"mycode/facility"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", ShowResponse)
	http.ListenAndServe(":"+port, nil)
}

func ShowResponse(w http.ResponseWriter, r *http.Request) {
	apiKey := "ed4I0nvodHiOPLphsGdci3PoK"
	res, err := http.Get("https://data.sfgov.org/resource/rqzj-sfat.json?$$app_token=" + apiKey)
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
	}
	data := make([]facility.Facility, 0)
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
