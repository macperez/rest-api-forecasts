package controllers

import (
	"encoding/json"
	"gitlab/moxdata/forecast-api/models"
	"log"
	"net/http"
	"time"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetForecastsFor(w http.ResponseWriter, r *http.Request) {
	const shortForm = "2006-01-02"
	start_date, ok := r.URL.Query()["start-date"]
	if !ok || len(start_date[0]) < 1 {
		log.Println("Url Param 'start-date' is missing")
		return
	}
	end_date, ok := r.URL.Query()["end-date"]
	if !ok || len(end_date[0]) < 1 {
		log.Println("Url Param 'end-date' is missing")
		return
	}

	city, ok := r.URL.Query()["city"]
	if !ok || len(city[0]) < 1 {
		log.Println("Url Param 'city' is missing")
		return
	}
	startDate, _ := time.Parse(shortForm, start_date[0])
	endDate, _ := time.Parse(shortForm, end_date[0])

	if startDate.After(endDate) {
		resp := Message(false, "Start date cannot be grater than end date")
		Respond(w, resp)
	} else {

		data := models.GetForecasts(start_date[0], end_date[0], city[0])
		resp := Message(true, "success")
		resp["forecasts"] = data
		Respond(w, resp)
	}

}
