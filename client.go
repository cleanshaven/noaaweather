package noaaweather

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"
)

func GetWeather(latitude, longitude string) (forecast ForecastJson, alerts AlertJson, err error) {
	forecast = ForecastJson{}
	alerts = AlertJson{}
	err = nil

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return
	}

	defer conn.Close()

	var forecastJsonString, alertJsonString string

	obj := conn.Object("com.github.cleanshaven.WeatherService", "/com/github/cleanshaven/WeatherService")
	err = obj.Call("com.github.cleanshaven.WeatherService.GetWeather", 0, longitude, latitude).Store(&forecastJsonString, &alertJsonString)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(forecastJsonString), &forecast)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(alertJsonString), &alerts)
	return

}
