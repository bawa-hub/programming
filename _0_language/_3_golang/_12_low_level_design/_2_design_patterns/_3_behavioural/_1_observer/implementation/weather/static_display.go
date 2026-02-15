package weather

import "fmt"

type StatisticsDisplay struct {
	id string
}

func NewStatisticsDisplay(id string) *StatisticsDisplay {
	return &StatisticsDisplay{id: id}
}

func (sd *StatisticsDisplay) Update(weather *WeatherData) {
	fmt.Printf("Statistics Display %s: %s\n", sd.id, weather)
	
	// Simple statistics
	if weather.Temperature > 30 {
		fmt.Printf("Statistics Display %s: Hot weather warning!\n", sd.id)
	} else if weather.Temperature < 0 {
		fmt.Printf("Statistics Display %s: Freezing weather warning!\n", sd.id)
	}
}

func (sd *StatisticsDisplay) GetID() string {
	return sd.id
}
