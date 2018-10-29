package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

var startTime = time.Now()
var dbStatus = ""

func main() {

	dbconnection, err := DBConnect()


	if dbconnection == true {
		dbStatus = "Connected to database"
	} else {
		dbStatus =  err.Error()
	}

	root := "/paragliding"
	r := mux.NewRouter()
	route := r.PathPrefix("/paragliding/api").Subrouter()

	http.HandleFunc(root, rootHandler)
	http.HandleFunc(root+"/api", apiHandler)
	http.HandleFunc(root+"/api/track", TrackHandler)
	route.HandleFunc("/track/{id}", GetTrackWithId).Methods("GET")

	http.Handle("/", r)

	appengine.Main()
}
/**
Function redirects the root url("https://igc-info-a2.herokuapp.com/paragliding)
to the apt url(https://igc-info-a2.herokuapp.com/paragliding/api)
 */
func rootHandler (w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/paragliding/api", 301)
}

/**

 */
func apiHandler(w http.ResponseWriter, r *http.Request) {

	type Info struct {
		Uptime string `json:"uptime"`
		Info string `json:"info"`
		Version string `json:"version"`
	}

	info := &Info{
		Uptime: uptime(),
		Info: "Service for Paragliding tracks.",
		Version: "v1",
	}

	i, _ := json.MarshalIndent(info, "", " ")

	fmt.Fprint(w, dbStatus, "\n\n", string(i))
}

/**
*	calculates time since boot un UNIX time and converts it into a Time object,
	afterwards converts it to a final string with all components
*/
func uptime() string{

	uptime := time.Now().Unix() - startTime.Unix()
	ut := time.Unix(uptime, 0)

	years, months, days 	:= ut.Date()
	hours, minutes, seconds := ut.Clock()
	months, days = months-1, days-1
	/**
	Return in ISO 8601 Format using P, Y, M, W, D, T, H, M, and S
	 */
	return fmt.Sprintf("%s%d%s%d%s%d%s%d%s%d%s%d%s", "P", absInt(int64(years - 1970)), "Y", months, "M", days, "DT", hours, "H", minutes, "M", seconds, "S")
}

/**
absolute value of n
  */
func absInt(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}