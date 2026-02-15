package weather

type WeatherSubject interface {
	Attach(observer WeatherObserver)
	Detach(observer WeatherObserver)
	Notify()
	SetMeasurements(temperature, humidity, pressure float64)
}