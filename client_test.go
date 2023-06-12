package noaaweather

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetWeather(t *testing.T) {
	forecast, alert, err := GetWeather("41.26", "-81.86")
	if err != nil {
		t.Fatalf("error from server %v", err)
	}

	spew.Dump(forecast)
	spew.Dump(alert)

}
