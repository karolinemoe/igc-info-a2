package main

import (
	"encoding/json"
	"fmt"
	"github.com/marni/goigc"
	"github.com/mitchellh/hashstructure"
	"google.golang.org/appengine"
	"net/http"
	"time"
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
	http.HandleFunc(root+"/api", apiHandler)
	http.HandleFunc(root+"/api/igc", igcHandler)

	appengine.Main()
}

/**

 */
func apiHandler(w http.ResponseWriter, r *http.Request) {




	fmt.Println("testttttt")

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

/**

 */
 func igcHandler(w http.ResponseWriter, r *http.Request) {
	 switch r.Method {
	 /**
	 If the http request method == GET
	  */
	 case http.MethodGet:
	 /**
	 If the http request method == POST
	  */
	 case http.MethodPost:
	 	var body struct{ URL string }
	 	err := json.NewDecoder(r.Body).Decode(&body)
	 	if err != nil {
	 		http.Error(w, err.Error(), 400); return
		}

	 	if body.URL == "" {
	 		http.Error(w, "Body of Request does not have a 'URL' property", 400); return
		}

	 	newTrack(body.URL, w)
	 default:
		 http.Error(w, "No specified request method", 400); return
	 }
 }

 /**

  */
 func newTrack(url string, w http.ResponseWriter) {
 	igcData, err := igc.ParseLocation(url)

 	if err != nil {
 		http.Error(w, "Problem parsing the URL", 400); return
	}

 	checksum, err := hashstructure.Hash(igcData, nil)
 	if err != nil {
 		http.Error(w, "Problem generating the hashstructure", 400); return
	}

 	trackID := int(checksum)

 	if !trackExists(trackID) {
 		// Save this track
	}
}

 func trackExists(trackID int) bool {
 	if (trackID == 1) {
 		return true
	} else {
		return false
 	}
 }