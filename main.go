package main

import (
	"runtime"
	"net/http"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	apiKey := ""
	response, err := http.Get("https://data.sfgov.org/resource/rqzj-sfat.json&$$app_token=" + apiKey)

	if err != nil {
		return
	}

	defer response.Body.Close()
}
