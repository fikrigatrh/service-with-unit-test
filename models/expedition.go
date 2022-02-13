package models

type ExpeditionSchedule struct {
	Vessel  string `json:"vessel"`
	Route   string `json:"route"`
	Etd     string `json:"etd"`
	Eta     string `json:"eta"`
	Closing string `json:"closing"`
}
