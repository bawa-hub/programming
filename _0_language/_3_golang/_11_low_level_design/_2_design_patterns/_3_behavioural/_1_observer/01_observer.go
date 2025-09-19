package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// BASIC OBSERVER PATTERN
// =============================================================================

// Observer interface
type Observer interface {
	Update(data interface{})
	GetID() string
}

// Subject interface
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
	GetState() interface{}
	SetState(state interface{})
}

// Concrete Subject
type ConcreteSubject struct {
	observers []Observer
	state     interface{}
	mu        sync.RWMutex
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
		state:     nil,
	}
}

func (cs *ConcreteSubject) Attach(observer Observer) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.observers = append(cs.observers, observer)
	fmt.Printf("Observer %s attached\n", observer.GetID())
}

func (cs *ConcreteSubject) Detach(observer Observer) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	
	for i, obs := range cs.observers {
		if obs.GetID() == observer.GetID() {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			fmt.Printf("Observer %s detached\n", observer.GetID())
			break
		}
	}
}

func (cs *ConcreteSubject) Notify() {
	cs.mu.RLock()
	observers := make([]Observer, len(cs.observers))
	copy(observers, cs.observers)
	state := cs.state
	cs.mu.RUnlock()
	
	fmt.Printf("Notifying %d observers\n", len(observers))
	for _, observer := range observers {
		observer.Update(state)
	}
}

func (cs *ConcreteSubject) GetState() interface{} {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.state
}

func (cs *ConcreteSubject) SetState(state interface{}) {
	cs.mu.Lock()
	cs.state = state
	cs.mu.Unlock()
	fmt.Printf("Subject state changed to: %v\n", state)
	cs.Notify()
}

// Concrete Observer
type ConcreteObserver struct {
	id string
}

func NewConcreteObserver(id string) *ConcreteObserver {
	return &ConcreteObserver{id: id}
}

func (co *ConcreteObserver) Update(data interface{}) {
	fmt.Printf("Observer %s received update: %v\n", co.id, data)
}

func (co *ConcreteObserver) GetID() string {
	return co.id
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. STOCK PRICE MONITORING SYSTEM
type StockPrice struct {
	Symbol    string
	Price     float64
	Change    float64
	ChangePercent float64
	Timestamp time.Time
}

func (sp *StockPrice) String() string {
	return fmt.Sprintf("%s: $%.2f (%.2f, %.2f%%)", 
		sp.Symbol, sp.Price, sp.Change, sp.ChangePercent)
}

type StockObserver interface {
	Update(stock *StockPrice)
	GetID() string
}

type StockSubject interface {
	Attach(observer StockObserver)
	Detach(observer StockObserver)
	Notify()
	UpdatePrice(symbol string, price float64)
	GetPrice(symbol string) *StockPrice
}

type StockMarket struct {
	observers []StockObserver
	stocks    map[string]*StockPrice
	mu        sync.RWMutex
}

func NewStockMarket() *StockMarket {
	return &StockMarket{
		observers: make([]StockObserver, 0),
		stocks:    make(map[string]*StockPrice),
	}
}

func (sm *StockMarket) Attach(observer StockObserver) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.observers = append(sm.observers, observer)
	fmt.Printf("Stock observer %s attached\n", observer.GetID())
}

func (sm *StockMarket) Detach(observer StockObserver) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	for i, obs := range sm.observers {
		if obs.GetID() == observer.GetID() {
			sm.observers = append(sm.observers[:i], sm.observers[i+1:]...)
			fmt.Printf("Stock observer %s detached\n", observer.GetID())
			break
		}
	}
}

func (sm *StockMarket) Notify() {
	sm.mu.RLock()
	observers := make([]StockObserver, len(sm.observers))
	copy(observers, sm.observers)
	sm.mu.RUnlock()
	
	for _, observer := range observers {
		// Notify about all stocks
		for _, stock := range sm.stocks {
			observer.Update(stock)
		}
	}
}

func (sm *StockMarket) UpdatePrice(symbol string, price float64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if existingStock, exists := sm.stocks[symbol]; exists {
		oldPrice := existingStock.Price
		existingStock.Price = price
		existingStock.Change = price - oldPrice
		existingStock.ChangePercent = (existingStock.Change / oldPrice) * 100
		existingStock.Timestamp = time.Now()
	} else {
		sm.stocks[symbol] = &StockPrice{
			Symbol:        symbol,
			Price:         price,
			Change:        0,
			ChangePercent: 0,
			Timestamp:     time.Now(),
		}
	}
	
	fmt.Printf("Stock price updated: %s\n", sm.stocks[symbol])
	sm.mu.Unlock()
	sm.Notify()
	sm.mu.Lock()
}

func (sm *StockMarket) GetPrice(symbol string) *StockPrice {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.stocks[symbol]
}

// Stock Observers
type StockTrader struct {
	id        string
	portfolio map[string]int
}

func NewStockTrader(id string) *StockTrader {
	return &StockTrader{
		id:        id,
		portfolio: make(map[string]int),
	}
}

func (st *StockTrader) Update(stock *StockPrice) {
	fmt.Printf("Trader %s: %s\n", st.id, stock)
	
	// Simple trading logic
	if stock.ChangePercent > 5.0 {
		fmt.Printf("Trader %s: BUY signal for %s (%.2f%% increase)\n", 
			st.id, stock.Symbol, stock.ChangePercent)
	} else if stock.ChangePercent < -5.0 {
		fmt.Printf("Trader %s: SELL signal for %s (%.2f%% decrease)\n", 
			st.id, stock.Symbol, stock.ChangePercent)
	}
}

func (st *StockTrader) GetID() string {
	return st.id
}

type StockAnalyst struct {
	id string
}

func NewStockAnalyst(id string) *StockAnalyst {
	return &StockAnalyst{id: id}
}

func (sa *StockAnalyst) Update(stock *StockPrice) {
	fmt.Printf("Analyst %s: Analyzing %s\n", sa.id, stock)
	
	// Simple analysis logic
	if stock.ChangePercent > 0 {
		fmt.Printf("Analyst %s: %s is performing well (%.2f%% up)\n", 
			sa.id, stock.Symbol, stock.ChangePercent)
	} else {
		fmt.Printf("Analyst %s: %s is underperforming (%.2f%% down)\n", 
			sa.id, stock.Symbol, stock.ChangePercent)
	}
}

func (sa *StockAnalyst) GetID() string {
	return sa.id
}

// 2. WEATHER STATION SYSTEM
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

type WeatherObserver interface {
	Update(weather *WeatherData)
	GetID() string
}

type WeatherSubject interface {
	Attach(observer WeatherObserver)
	Detach(observer WeatherObserver)
	Notify()
	SetMeasurements(temperature, humidity, pressure float64)
}

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

// Weather Observers
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

// 3. EVENT SYSTEM
type Event struct {
	Type      string
	Data      interface{}
	Timestamp time.Time
	Source    string
}

func (e *Event) String() string {
	return fmt.Sprintf("[%s] %s from %s: %v", 
		e.Timestamp.Format("15:04:05"), e.Type, e.Source, e.Data)
}

type EventObserver interface {
	HandleEvent(event *Event)
	GetID() string
	GetEventTypes() []string
}

type EventSubject interface {
	Subscribe(observer EventObserver)
	Unsubscribe(observer EventObserver)
	Publish(event *Event)
}

type EventBus struct {
	observers map[string][]EventObserver
	mu        sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		observers: make(map[string][]EventObserver),
	}
}

func (eb *EventBus) Subscribe(observer EventObserver) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	for _, eventType := range observer.GetEventTypes() {
		eb.observers[eventType] = append(eb.observers[eventType], observer)
	}
	fmt.Printf("Event observer %s subscribed to events\n", observer.GetID())
}

func (eb *EventBus) Unsubscribe(observer EventObserver) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	
	for eventType, observers := range eb.observers {
		for i, obs := range observers {
			if obs.GetID() == observer.GetID() {
				eb.observers[eventType] = append(observers[:i], observers[i+1:]...)
				break
			}
		}
	}
	fmt.Printf("Event observer %s unsubscribed from events\n", observer.GetID())
}

func (eb *EventBus) Publish(event *Event) {
	eb.mu.RLock()
	observers := eb.observers[event.Type]
	eb.mu.RUnlock()
	
	fmt.Printf("Publishing event: %s\n", event)
	for _, observer := range observers {
		observer.HandleEvent(event)
	}
}

// Event Observers
type LoggingObserver struct {
	id string
}

func NewLoggingObserver(id string) *LoggingObserver {
	return &LoggingObserver{id: id}
}

func (lo *LoggingObserver) HandleEvent(event *Event) {
	fmt.Printf("Logging Observer %s: Logging event - %s\n", lo.id, event)
}

func (lo *LoggingObserver) GetID() string {
	return lo.id
}

func (lo *LoggingObserver) GetEventTypes() []string {
	return []string{"user_login", "user_logout", "error", "warning"}
}

type NotificationObserver struct {
	id string
}

func NewNotificationObserver(id string) *NotificationObserver {
	return &NotificationObserver{id: id}
}

func (no *NotificationObserver) HandleEvent(event *Event) {
	fmt.Printf("Notification Observer %s: Sending notification for - %s\n", no.id, event)
}

func (no *NotificationObserver) GetID() string {
	return no.id
}

func (no *NotificationObserver) GetEventTypes() []string {
	return []string{"user_login", "error", "warning", "system_alert"}
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== OBSERVER PATTERN DEMONSTRATION ===\n")

	// 1. BASIC OBSERVER
	fmt.Println("1. BASIC OBSERVER:")
	subject := NewConcreteSubject()
	
	observer1 := NewConcreteObserver("Observer1")
	observer2 := NewConcreteObserver("Observer2")
	observer3 := NewConcreteObserver("Observer3")
	
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.SetState("Hello, World!")
	
	subject.Attach(observer3)
	subject.SetState("Updated state")
	
	subject.Detach(observer2)
	subject.SetState("Final state")
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Stock Price Monitoring
	fmt.Println("Stock Price Monitoring System:")
	stockMarket := NewStockMarket()
	
	trader1 := NewStockTrader("Trader1")
	trader2 := NewStockTrader("Trader2")
	analyst1 := NewStockAnalyst("Analyst1")
	
	stockMarket.Attach(trader1)
	stockMarket.Attach(trader2)
	stockMarket.Attach(analyst1)
	
	// Update stock prices
	stockMarket.UpdatePrice("AAPL", 150.0)
	stockMarket.UpdatePrice("GOOGL", 2800.0)
	stockMarket.UpdatePrice("AAPL", 155.0) // 3.33% increase
	stockMarket.UpdatePrice("GOOGL", 2700.0) // -3.57% decrease
	
	stockMarket.Detach(trader2)
	stockMarket.UpdatePrice("MSFT", 300.0)
	fmt.Println()

	// Weather Station
	fmt.Println("Weather Station System:")
	weatherStation := NewWeatherStation()
	
	currentDisplay := NewCurrentConditionsDisplay("Display1")
	statsDisplay := NewStatisticsDisplay("Stats1")
	forecastDisplay := NewForecastDisplay("Forecast1")
	
	weatherStation.Attach(currentDisplay)
	weatherStation.Attach(statsDisplay)
	weatherStation.Attach(forecastDisplay)
	
	// Update weather measurements
	weatherStation.SetMeasurements(25.0, 65.0, 1013.0)
	weatherStation.SetMeasurements(32.0, 70.0, 1005.0)
	weatherStation.SetMeasurements(-5.0, 80.0, 1020.0)
	fmt.Println()

	// Event System
	fmt.Println("Event System:")
	eventBus := NewEventBus()
	
	loggingObserver := NewLoggingObserver("Logger1")
	notificationObserver := NewNotificationObserver("Notifier1")
	
	eventBus.Subscribe(loggingObserver)
	eventBus.Subscribe(notificationObserver)
	
	// Publish events
	eventBus.Publish(&Event{
		Type:      "user_login",
		Data:      "User John logged in",
		Timestamp: time.Now(),
		Source:    "auth_service",
	})
	
	eventBus.Publish(&Event{
		Type:      "error",
		Data:      "Database connection failed",
		Timestamp: time.Now(),
		Source:    "database_service",
	})
	
	eventBus.Publish(&Event{
		Type:      "warning",
		Data:      "High memory usage detected",
		Timestamp: time.Now(),
		Source:    "monitoring_service",
	})
	
	eventBus.Publish(&Event{
		Type:      "system_alert",
		Data:      "Server maintenance scheduled",
		Timestamp: time.Now(),
		Source:    "admin_service",
	})
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
