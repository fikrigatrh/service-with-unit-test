package models

import "gorm.io/gorm"

type ExpeditionSchedule struct {
	gorm.Model
	Vessel  string `json:"vessel"`
	Route   string `json:"route"`
	Etd     string `json:"etd"`
	Eta     string `json:"eta"`
	Closing string `json:"closing"`
}

type PriceList struct {
	Dari               string `json:"dari"`
	Tujuan             string `json:"tujuan"`
	HargaPerKilo       string `json:"harga_per_kilo"`
	HargaPerMeterKubik string `json:"harga_per_meter_kubik"`
}
