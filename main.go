package main

import (
	"Excercise/Crawler"
	"Excercise/DBEngine"
	"Excercise/model"
)

func main() {
	db := DBEngine.InitDBEngine()
	db.AutoMigrate(&model.Areastate{})
	response := Crawler.Crawler()
	db.Create(&response)
}
