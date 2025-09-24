package weather

type WeatherObserver interface {
	Update(weather *WeatherData)
	GetID() string
}