package noaaweather

import (
	"encoding/json"

	"github.com/godbus/dbus/v5"
)

const (
	ServerName     = "com.github.cleanshaven.WeatherService"
	ServerPath     = "/com/github/cleanshaven/WeatherService"
	GetWeatherCall = ServerName + ".GetWeather"
	GetIconCall    = ServerName + ".GetIcon"
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

	obj := conn.Object(ServerName, ServerPath)
	err = obj.Call(GetWeatherCall, 0, longitude, latitude).Store(&forecastJsonString, &alertJsonString)
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

func GetIcon(iconUrl string) (filename string, err error) {
	filename = ""
	err = nil

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return
	}

	defer conn.Close()

	obj := conn.Object(ServerName, ServerPath)
	err = obj.Call(GetIconCall, 0, iconUrl).Store(&filename)
	return
}
