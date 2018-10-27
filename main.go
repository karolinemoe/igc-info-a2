package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
)

func main() {

	root := "/paragliding"
	http.HandleFunc(root+"/api", apiHandler)

	appengine.Main()
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	type Info struct {
		Uptime string `json:"uptime"`
		Info string `json:"info"`
		Version string `json:"version"`
	}

	info := &Info{
		//Uptime: uptime(),
		Info: "Service for IGC tracks.",
		Version: "v1",
	}

	i, _ := json.MarshalIndent(info, "", " ")

	fmt.Fprint(w, string(i))
}
