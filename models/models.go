package models

import (
	"time"
)

const jsonTimeLayout = "2006-01-02T15:04:05"

type ForecastRecord struct {
	Partner  string     `json:"Partner"`
	Id       int        `json:"Id"`
	Name     string     `json:"Name"`
	Time     *time.Time `json:"Time"`
	Forecast float32    `gorm:"type:float" json:"Forecast"`
}

type Forecast struct {
	Partner string `json:"Partner"`
	Hubs    []Hub  `json:"hub"`
}

type ForecastItem struct {
	Time     string  `json:"Time"`
	Forecast float32 `gorm:"type:float" json:"Forecast"`
}

type Hub struct {
	Id            int            `json:"Id"`
	Name          string         `json:"Name"`
	ForecastItems []ForecastItem `json:"forecasts"`
}

func GetForecasts(startDate, endDate, city string) []Forecast {
	sql_query := `
	SELECT "Burger King" AS partner, r.restaurant_id as id, r.name AS name, 
	  f.bk_datetime AS time, f.hours AS forecast
	FROM forecast AS f
	  JOIN restaurant AS r USING (restaurant_id)
	  JOIN miscellany.city AS ct ON (ct.id = r.city)
	WHERE DATE(bk_datetime) BETWEEN DATE(?) AND DATE(?)
	  AND ct.name = ? 
	;`

	forecasts := make([]Forecast, 0)
	forecastsRecords := make([]*ForecastRecord, 0)
	GetDB().Raw(sql_query, startDate, endDate, city).Scan(&forecastsRecords)

	forecast := Forecast{
		Partner: "Burger King",
	}
	forecast.Hubs = make([]Hub, 0)
	hub := Hub{
		Id:   -1,
		Name: "",
	}
	var forecast_items []ForecastItem
	for _, forecast_record := range forecastsRecords {
		if forecast_record.Id != hub.Id {
			hub.ForecastItems = forecast_items
			if hub.Id >= 0 {
				forecast.Hubs = append(forecast.Hubs, hub)
			}
			hub = Hub{
				Id:   forecast_record.Id,
				Name: forecast_record.Name,
			}

			f_item := ForecastItem{
				Time:     forecast_record.Time.Format(jsonTimeLayout),
				Forecast: forecast_record.Forecast,
			}
			forecast_items = make([]ForecastItem, 0)
			forecast_items = append(forecast_items, f_item)

		} else {

			f_item := ForecastItem{
				Time:     forecast_record.Time.Format(jsonTimeLayout),
				Forecast: forecast_record.Forecast,
			}
			forecast_items = append(forecast_items, f_item)
		}

	}

	forecasts = append(forecasts, forecast)
	return forecasts
}
