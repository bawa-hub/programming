package weather

import "fmt"

type CurrentConditionsDisplay struct {
	id string
}

func NewCurrentConditionsDisplay(id string) *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{id: id}
}

func (ccd *CurrentConditionsDisplay) Update(weather *WeatherData) {
	fmt.Printf("Current Conditions Display %s: %s\n", ccd.id, weather)
}

func (ccd *CurrentConditionsDisplay) GetID() string {
	return ccd.id
}