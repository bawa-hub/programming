package weather

import (
	"fmt"
	"time"
)

type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	Timestamp   time.Time
}

func (wd *WeatherData) String() string {
	return fmt.Sprintf("Temp: %.1f°C, Humidity: %.1f%%, Pressure: %.1f hPa", 
		wd.Temperature, wd.Humidity, wd.Pressure)
}
