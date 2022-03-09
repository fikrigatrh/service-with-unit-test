package models

import "gorm.io/gorm"

type ExpeditionSchedule struct {
	gorm.Model
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
