package main

import (
	"os"
	"log"
	"fmt"
	"runtime"
	"net/http"
	"encoding/json"
)

type Location struct {
	needs_recoding	bool		`json:"needs_recoding"`
	longitude				string	`json:"longitude"`
	latitude				string	`json:"latitude"`
}

type Facility struct {
	Location						*Location
	Status							string	`json:"status,omitempty"`
	Expirationdate			string	`json:"expirationdate,omitempty"`
	Permit							string	`json:"permit,omitempty"`
	Block								string	`json:"block,omitempty"`
	Received						string	`json:"received,omitempty"`
	Facilitytype				string	`json:"facilitytype,omitempty"`
	Blocklot						string	`json:"blocklot,omitempty"`
	Noisent							string	`json:"noisent,omitempty"`
	Locationdescription	string	`json:"locationdescription,omitempty"`
	Cnn									string	`json:"cnn,omitempty"`
	Priorpermit					string	`json:"priorpermit,omitempty"`
	Approved						string	`json:"approved,omitempty"`
	Schedule						string	`json:"schedule,omitempty"`
	Address							string	`json:"address,omitempty"`
	Applicant						string	`json:"applicant,omitempty"`
	Lot									string	`json:"lot,omitempty"`
	Fooditems						string	`json:"fooditems,omitempty"`
	Longitude						string	`json:"longitude,omitempty"`
	Latitude						string	`json:"latitude,omitempty"`
	Objectid						string	`json:"objectid,omitempty"`
	Y										string	`json:"y,omitempty"`
	X										string	`json:"x,omitempty"`
}

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
	data := make([]Facility, 0)
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	fmt.Println(&data)
}
