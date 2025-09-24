package weather

import (
	"fmt"
	"sync"
	"time"
)

type WeatherStation struct {
	observers []WeatherObserver
	weather   *WeatherData
	mu        sync.RWMutex
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{
		observers: make([]WeatherObserver, 0),
		weather:   &WeatherData{},
	}
}

func (ws *WeatherStation) Attach(observer WeatherObserver) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.observers = append(ws.observers, observer)
	fmt.Printf("Weather observer %s attached\n", observer.GetID())
}

func (ws *WeatherStation) Detach(observer WeatherObserver) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	
	for i, obs := range ws.observers {
		if obs.GetID() == observer.GetID() {
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
			fmt.Printf("Weather observer %s detached\n", observer.GetID())
			break
		}
	}
}

func (ws *WeatherStation) Notify() {
	ws.mu.RLock()
	observers := make([]WeatherObserver, len(ws.observers))
	copy(observers, ws.observers)
	weather := ws.weather
	ws.mu.RUnlock()
	
	for _, observer := range observers {
		observer.Update(weather)
	}
}

func (ws *WeatherStation) SetMeasurements(temperature, humidity, pressure float64) {
	ws.mu.Lock()
	ws.weather.Temperature = temperature
	ws.weather.Humidity = humidity
	ws.weather.Pressure = pressure
	ws.weather.Timestamp = time.Now()
	ws.mu.Unlock()
	
	fmt.Printf("Weather measurements updated: %s\n", ws.weather)
	ws.Notify()
}