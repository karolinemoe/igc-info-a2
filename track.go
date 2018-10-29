package main


import (
	"encoding/json"
	"github.com/marni/goigc"
	"github.com/mitchellh/hashstructure"
	"net/http"
	"strconv"
	"time"
)

// IGCTrack struct for the track data
type IGCTrack struct {
	HDate       time.Time `json:"H_date"`
	Pilot       string    `json:"pilot"`
	Glider      string    `json:"glider"`
	GliderID    string    `json:"glider_id"`
	ID          string    `json:"track_id"`
	TrackLength float64   `json:"track_length"`
}

/**
 */
func TrackHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	/**
	If the http request method == GET
	 */
	case http.MethodGet:
		tracks := GetTracks()
		json.NewEncoder(w).Encode(tracks); return

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

		newID := newTrack(body.URL, w)
		if newID < 0 {
			http.Error(w, "Not able to process the URL", http.StatusBadRequest); return
		}
	default:
		http.Error(w, "No specified request method", 400); return
	}
}

/**
 */
func newTrack(url string, w http.ResponseWriter) int {
	igcData, err := igc.ParseLocation(url)

	if err != nil {
		http.Error(w, "Problem parsing the URL", 400)
	}

	checksum, err := hashstructure.Hash(igcData, nil)
	if err != nil {
		http.Error(w, "Problem generating the hashstructure", 400)
	}

	trackID := int(checksum)

	/**
	store data in memory
	  */
	trackData := IGCTrack{
		HDate:       igcData.Date,
		Pilot:       igcData.Pilot,
		Glider:      igcData.GliderType,
		GliderID:    igcData.GliderID,
		TrackLength: calcTrackLength(igcData.Points),
		ID:     	strconv.Itoa(int(trackID)),
	}

	InsertTrack(trackData)

	type IGCid struct {
		ID int `json:"id"`
	}

	json.NewEncoder(w).Encode(IGCid{ID: trackID})
	return trackID
}

func calcTrackLength(points []igc.Point) float64 {
	var tl float64
	for i := 0; i < len(points)-1; i++ {
		tl += points[i].Distance(points[i+1])
	}
	return tl
}
