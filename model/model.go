package model

import "github.com/jinzhu/gorm"

type Roadstate struct {
	gorm.Model
	Name      string `json:"name"`
	Status    string `json:"status"`
	Direction string `json:"direction"`
	Angle     string `json:"angle"`
	Speed     string `json:"speed"`
	Polyline  string `json:""polyline`
}
type Areastate struct {
	gorm.Model
	Expedite    string `json:"expedite"`
	Congested   string `json:"congested"`
	Blocked     string `json:"blocked"`
	Unknown     string `json:"unknown"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
type Trafficinfo struct {
	Description string    `json:"description"`
	Areastate   Areastate `json:"evaluation"`
	Roadstate   Roadstate `json:"roads"`
}
type RespenseStatus struct {
	Status      string      `json:"status"`
	Info        string      `json:"info"`
	InfoCode    string      `json:"infocode"`
	Trafficinfo Trafficinfo `json:"trafficinfo"`
}
