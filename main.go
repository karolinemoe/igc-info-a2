package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
	"time"
)

var startTime = time.Now()

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
		Uptime: uptime(),
		Info: "Service for Paragliding tracks.",
		Version: "v1",
	}

	i, _ := json.MarshalIndent(info, "", " ")

	fmt.Fprint(w, string(i))
}

// calculates time since boot un UNIX time and converts it into a Time object, afterwards converst it to a final string with all components
func uptime() string{

	uptime := time.Now().Unix() - startTime.Unix()
	ut := time.Unix(uptime, 0)

	years, months, days 	:= ut.Date()
	hours, minutes, seconds := ut.Clock()
	months, days = months-1, days-1

	return fmt.Sprintf("%s%d%s%d%s%d%s%d%s%d%s%d%s", "P", absInt(int64(years - 1970)), "Y", months, "M", days, "DT", hours, "H", minutes, "M", seconds, "S")
}

// absolute value of n
func absInt(n int64) int64 {

	if n < 0 {
		return -n
	}
	return n
}
