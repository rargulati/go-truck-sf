package main

import (
	"os"
	"log"
	"fmt"
	"runtime"
	"net/http"
	"encoding/json"
)

var APIKey string
/*
type location struct {
	needs_recoding	bool		`json:"needs_recoding"`
	longitude				string	`json:"longitude"`
	latitude				string	`json:"latitude"`
}

type mobileFood struct {
	location						json.RawMessage
	status							string	`json:"status,omitempty"`
	expirationdate			string	`json:"expirationdate,omitempty"`
	permit							string	`json:"permit,omitempty"`
	block								string	`json:"block,omitempty"`
	received						string	`json:"received,omitempty"`
	facilitytype				string	`json:"facilitytype,omitempty"`
	blocklot						string	`json:"blocklot,omitempty"`
	noisent							string	`json:"noisent,omitempty"`
	locationdescription	string	`json:"locationdescription,omitempty"`
	cnn									string	`json:"cnn,omitempty"`
	priorpermit					string	`json:"priorpermit,omitempty"`
	approved						string	`json:"approved,omitempty"`
	schedule						string	`json:"schedule,omitempty"`
	address							string	`json:"address,omitempty"`
	applicant						string	`json:"applicant,omitempty"`
	lot									string	`json:"lot,omitempty"`
	fooditems						string	`json:"fooditems,omitempty"`
	longitude						string	`json:"longitude,omitempty"`
	latitude						string	`json:"latitude,omitempty"`
	objectid						string	`json:"objectid,omitempty"`
	y										string	`json:"y,omitempty"`
	x										string	`json:"x,omitempty"`
}
*/

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
	var data interface{}
	defer res.Body.Close()
	// w.Header().Set("Content_type", "application/json")
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	m := data.([]interface{})
	for _,info := range m {
		r := info.(map[string]interface{})
		for k,v := range r {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			case map[string]interface{}:
				fmt.Println(k, "is nested json:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle", v)
			}
		}
	}
}
