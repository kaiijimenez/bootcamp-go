package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kaiijimenez/bootcamp-go/earthquakes/models"
)

func RespondWithError(e error, code int, message string, w http.ResponseWriter) {
	var err models.ErrorResponse
	log.Fatal(message, e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err.Code = code
	err.Msg = message
	json.NewEncoder(w).Encode(&err)
}

func RespondWithJson(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&res)
}

func ConvertToRes(js models.SpecificID) models.Response {
	var res models.Response
	res.Title = js.Propertie.Title
	res.Time = fmt.Sprintf("%v", time.Unix(js.Propertie.OTime, 0))
	res.Lat = GetLat(js.Geo.Coordinates[1])
	res.Lon = GetLon(js.Geo.Coordinates[0])
	res.Depth = fmt.Sprintf("%.2f km depth", js.Geo.Coordinates[2])
	res.Type = js.Propertie.Type
	res.Magnitud = fmt.Sprintf("%v ml", js.Propertie.Magnitud)
	res.ID = js.ID
	res.UpdatedT = fmt.Sprintf("%v", time.Unix(js.Propertie.UTime, 0))
	return res
}

func CheckAvailability(key string, earthquakes []models.Response) bool {
	for _, v := range earthquakes {
		if v.ID != key {
			return false
		}
	}
	return true
}

func GetLon(lon float64) string {
	if 0 < lon && lon < 180 {
		return fmt.Sprintf("%.2f °E", lon)
	}
	return fmt.Sprintf("%.2f °W", lon)
}

func GetLat(lat float64) string {
	if lat > 0 {
		return fmt.Sprintf("%.2f °N", lat)
	}
	return fmt.Sprintf("%.2f °S", lat)
}
