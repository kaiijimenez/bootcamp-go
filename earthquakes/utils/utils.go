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

func ConvertToRes(js models.SpecificID, res models.Response) models.Response {
	res.Title = fmt.Sprintf("Summary: %s", js.Propertie.Title)
	res.Time = fmt.Sprintf("Origin time: %v", time.Unix(js.Propertie.OTime, 0))
	res.Lat = fmt.Sprintf("Lat: %0.2f", js.Geo.Coordinates[1])
	res.Lon = fmt.Sprintf("Lat: %0.2f", js.Geo.Coordinates[0])
	res.Depth = fmt.Sprintf("%0.2f km depth", js.Geo.Coordinates[2])
	res.Type = fmt.Sprintf("%s", js.Propertie.Type)
	res.Magnitud = fmt.Sprintf("Magnitude %v ml", js.Propertie.Magnitud)
	res.ID = fmt.Sprintf("ID: %v", js.ID)
	res.UpdatedT = fmt.Sprintf("Updated time: %v", time.Unix(js.Propertie.UTime, 0))
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
