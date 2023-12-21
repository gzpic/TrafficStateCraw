package Crawler

import (
	"Excercise/Configs"
	"Excercise/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const key = "ad92bae70e179f0ee276c49650e3abe7"

var sc = Configs.StatesConfigs{"113.295185", "23.134381", "500"}

func Crawler() model.Areastate {
	website := "https://restapi.amap.com/v3/traffic/status/circle?location=%s,%s&radius=%s&key=%s"
	url := fmt.Sprintf(website, sc.LocationX, sc.LocationY, sc.Radius, key)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Web err", err)
	}

	content, err := ioutil.ReadAll(resp.Body)
	conten := string(content)
	if err != nil {
		log.Fatalf("err", err)
	}
	var re model.RespenseStatus
	err = json.Unmarshal([]byte(conten), &re)
	if err != nil {
		log.Fatalf("ERR", err)
	}
	return re.Trafficinfo.Areastate
}
