package models

import "gorm.io/gorm"

type ResponseDataPagination struct {
	TotalData     int                  `json:"total_data"`
	DataPerPage   int                  `json:"data_per_page"`
	Page          int                  `json:"page"`
	TotalPage     int                  `json:"total_page"`
	NumberCurrent int                  `json:"number_current"`
	NumberEnd     int                  `json:"number_end"`
	NextUrlPage   string               `json:"next_url_page"`
	PrevUrlPage   string               `json:"prev_url_page"`
	Data          []ExpeditionSchedule `json:"data"`
}

type ExpeditionSchedule struct {
	gorm.Model
	No               int    `json:"no" gorm:"-"`
	Vessel           string `json:"vessel" validate:"required"`
	RouteFrom        string `json:"route_from,omitempty" validate:"required"`
	RouteDestination string `json:"route_destination,omitempty" validate:"required"`
	Route            string `json:"route"`
	Etd              string `json:"etd" validate:"required"`
	Eta              string `json:"eta" validate:"required"`
	Closing          string `json:"closing" validate:"required"`
}

type PriceList struct {
	Dari               string `json:"dari"`
	Tujuan             string `json:"tujuan"`
	HargaPerKilo       string `json:"harga_per_kilo"`
	HargaPerMeterKubik string `json:"harga_per_meter_kubik"`
}
