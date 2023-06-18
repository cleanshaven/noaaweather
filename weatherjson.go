package noaaweather

// https://api.weather.gov/alerts/active?point=41.26,-81.86
type GeometryJson struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type JsonNumbers interface {
	int | float64
}

type UnitValueJson[T JsonNumbers] struct {
	UnitCode string `json:"unitCode"`
	Value    T      `json:"value"`
}

type RelativeLocationPropertiesJson struct {
	City     string                 `json:"city"`
	State    string                 `json:"state"`
	Distance UnitValueJson[float64] `json:"distance"`
	Bearing  UnitValueJson[int]     `json:"bearing"`
}

type RelativeLocationJson struct {
	Type     string       `json:"type"`
	Geometry GeometryJson `json:"geometry"`
}

type PropertiesJson struct {
	Id                  string               `json:"@id"`
	Type                string               `json:"@type"`
	CWA                 string               `json:"cwa"`
	ForecastOffice      string               `json:"forecastOffice"`
	GridId              string               `json:"gridId"`
	GridX               int                  `json:"gridX"`
	GridY               int                  `json:"gridY"`
	Forecast            string               `json:"forecast"`
	ForecastHourly      string               `json:"forecastHourly"`
	ForcastGridData     string               `json:"forecastGridData"`
	ObservationStations string               `json:"observationStations"`
	RelativeLocation    RelativeLocationJson `json:"relativeLocation"`
	ForecastZone        string               `json:"forecastZone"`
	Country             string               `json:"country"`
	FireWeatherZone     string               `json:"fireWeatherZone"`
	TimeZone            string               `json:"timeZone"`
	RadarStation        string               `json:"radarStation"`
}

type LocationJson struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Geometry   GeometryJson   `json:"geometry"`
	Properties PropertiesJson `json:"properties"`
}

type PeriodJson struct {
	Number                     int                    `json:"number"`
	Name                       string                 `json:"name"`
	StartTime                  string                 `json:"startTime"`
	EndTime                    string                 `json:"endTime"`
	IsDaytime                  bool                   `json:"isDaytime"`
	Temperature                int                    `json:"temperature"`
	TemperatureUnit            string                 `json:"temperatureUnit"`
	TemperatureTrend           string                 `json:"temperatureTrend"`
	ProbabilityOfPrecipitation UnitValueJson[int]     `json:"probabilityOfPrecipitation"`
	DewPoint                   UnitValueJson[float64] `json:"dewpoint"`
	RelativeHumidity           UnitValueJson[int]     `json:"relativeHumidity"`
	WindSpeed                  string                 `json:"windSpeed"`
	WindDirection              string                 `json:"windDirection"`
	Icon                       string                 `json:"icon"`
	ShortForecast              string                 `json:"shortForecast"`
	DetailedForcast            string                 `json:"detailedForecast"`
}

type WeatherPropertiesJson struct {
	Updated          string                 `json:"updated"`
	Units            string                 `json:"units"`
	ForcastGenerator string                 `json:"forecastGenerator"`
	GeneratedAt      string                 `json:"generatedAt"`
	UpdateTime       string                 `json:"updateTime"`
	ValidTimes       string                 `json:"validTimes"`
	Elevation        UnitValueJson[float64] `json:"elevation"`
	Periods          []PeriodJson           `json:"periods"`
}

type ForecastJson struct {
	Type       string                `json:"type"`
	Properties WeatherPropertiesJson `json:"properties"`
}

type GeoCodeJson struct {
	SAME []string `json:"SAME"`
	UGC  []string `json:"UGC"`
}

type AlertPropertyJson struct {
	URLId         string      `json:"@id"`
	Type          string      `json:"@type"`
	Id            string      `json:"id"`
	AreaDesc      string      `json:"areaDesc"`
	GeoCode       GeoCodeJson `json:"geocode"`
	AffectedZones []string    `json:"affectedZones"`
	Sent          string      `json:"sent"`
	Effective     string      `json:"effective"`
	Onset         string      `json:"onset"`
	Expires       string      `json:"expires"`
	Ends          string      `json:"ends"`
	Status        string      `json:"status"`
	MessageType   string      `json:"messageType"`
	Category      string      `json:"category"`
	Severity      string      `json:"severity"`
	Certainty     string      `json:"certainty"`
	Urgency       string      `json:"urgency"`
	Event         string      `json:"event"`
	Sender        string      `json:"sender"`
	SenderName    string      `json:"senderName"`
	Headline      string      `json:"headline"`
	Description   string      `json:"description"`
	Instruction   string      `json:"instruction"`
	Response      string      `json:"response"`
}

type AlertFeatureJson struct {
	Id         string            `json:"id"`
	Type       string            `json:"type"`
	Properties AlertPropertyJson `json:"properties"`
}

type AlertJson struct {
	Type     string             `json:"type"`
	Features []AlertFeatureJson `json:"features"`
}
