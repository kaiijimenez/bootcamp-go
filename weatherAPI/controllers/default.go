package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type WeatherController struct {
	beego.Controller
}

type mystruct struct {
	Coord struct {
		Lon float64
		Lat float64
	} `json:"coord"`
}

func (wc *WeatherController) Get() {
	// city := wc.GetString("city")
	// country := wc.GetString("country")

	response := getResponse()
	coord := response["coord"].(map[string]interface{})
	fmt.Println(coord)
	wc.Data["json"] = &coord
	wc.ServeJSON()
}

func getResponse() map[string]interface{} {
	var jsonResponse map[string]interface{}
	res := httplib.Get("http://api.openweathermap.org/data/2.5/weather?q=Bogota,co&appid=1508a9a4840a5574c822d70ca2132032")
	str, _ := res.String()
	err := json.Unmarshal([]byte(str), &jsonResponse)
	if err != nil {
		fmt.Println(err)
	}
	return jsonResponse
}
