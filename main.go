package main

import (
	"fmt"
	"net/http"
)

func main() {

	root := "/paragliding"
	http.HandleFunc(root+"/api", apiHandler)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Application")
}
