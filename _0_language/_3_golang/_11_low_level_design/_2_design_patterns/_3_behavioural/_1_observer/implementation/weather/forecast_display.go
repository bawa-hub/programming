package weather

import "fmt"


type ForecastDisplay struct {
	id string
}

func NewForecastDisplay(id string) *ForecastDisplay {
	return &ForecastDisplay{id: id}
}

func (fd *ForecastDisplay) Update(weather *WeatherData) {
	fmt.Printf("Forecast Display %s: %s\n", fd.id, weather)
	
	// Simple forecast logic
	if weather.Pressure > 1013 {
		fmt.Printf("Forecast Display %s: High pressure - sunny weather expected\n", fd.id)
	} else if weather.Pressure < 1000 {
		fmt.Printf("Forecast Display %s: Low pressure - rainy weather expected\n", fd.id)
	}
}

func (fd *ForecastDisplay) GetID() string {
	return fd.id
}