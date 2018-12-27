package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kaiijimenez/bootcamp-go/earthquakes/utils"

	"github.com/kaiijimenez/bootcamp-go/earthquakes/models"

	"github.com/gorilla/mux"
)

const endepoint = "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&%s"

var earthquakes []models.SpecificID

func GetEarthquake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting earthquake by id")

	//http.Client with a sensible timeout
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	id := mux.Vars(r)["key"]
	eventid := fmt.Sprintf("eventid=%s", id)
	getE := fmt.Sprintf(endepoint, eventid)
	// if found 200 if not found 404 not found
	//Getting items from endpoint
	endp, _ := client.Get(getE)
	body, _ := ioutil.ReadAll(endp.Body)
	var fromID models.SpecificID
	err := json.Unmarshal(body, &fromID)
	if err != nil {
		utils.RespondWithError(err, http.StatusBadRequest, "Error trying to unsmarshal data", w)
	}
	earthquakes = append(earthquakes, fromID)
	utils.RespondWithJson(w, http.StatusOK, earthquakes)
}

func DeleteFromMemory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting by id")
	key := mux.Vars(r)["key"]
	//check if it is not in memory
	for index, v := range earthquakes {
		if v.ID == key {
			earthquakes = append(earthquakes[:index], earthquakes[index+1:]...)
			utils.RespondWithJson(w, http.StatusNoContent, earthquakes)
			return
		}
	}
	utils.RespondWithError(nil, http.StatusNotFound, "ID not found in memory", w)
}

func UpdateEarthquake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating by id")
	key := mux.Vars(r)["key"]
	var response models.SpecificID
	for i, v := range earthquakes {
		if v.ID == key {
			//deleting
			earthquakes = append(earthquakes[:i], earthquakes[i+1:]...)
			if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
				utils.RespondWithError(err, http.StatusBadRequest, "Error trying to decode the request", w)
				return
			}
			earthquakes = append(earthquakes, response)
			utils.RespondWithJson(w, http.StatusAccepted, earthquakes)
			return
		}
	}
	utils.RespondWithError(nil, http.StatusNotFound, "Error ID not found in the memory", w)
	return
}

func GetFromPeriod(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting from period of time")
	//var period models.PeriodTime
	/*period := make(map[string]string)
	fmt.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&period); err != nil {
		utils.RespondWithError(err, http.StatusBadRequest, "Error trying to decode the request ", w)
		return
	}*/
	timePeriod := fmt.Sprintf("starttime=%s&endtime=%s", "2018-12-19", "2018-12-20")
	getURL := fmt.Sprintf(endepoint, timePeriod)
	endp, _ := http.Get(getURL)
	body, _ := ioutil.ReadAll(endp.Body)
	var collection models.FeatureCollection
	if err := json.Unmarshal(body, &collection); err != nil {
		utils.RespondWithError(err, http.StatusNotImplemented, "Error trying to unmarshall the data", w)
		return
	}
	utils.RespondWithJson(w, http.StatusOK, collection)
	return
}

func ReportEarthquake(w http.ResponseWriter, r *http.Request) {
	var response models.SpecificID
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		utils.RespondWithError(err, http.StatusBadRequest, "Error trying to unmarshall the data", w)
		return
	}
	earthquakes = append(earthquakes, response)
	utils.RespondWithJson(w, http.StatusCreated, earthquakes)
	return
}
