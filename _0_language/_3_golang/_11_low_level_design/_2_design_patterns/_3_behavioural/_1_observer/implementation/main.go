package main

import (
	"fmt"
	"observer/stocks"
)


func main() {

	// // 1. BASIC OBSERVER
	// fmt.Println("1. BASIC OBSERVER:")
	// subject := NewConcreteSubject()
	
	// observer1 := NewConcreteObserver("Observer1")
	// observer2 := NewConcreteObserver("Observer2")
	// observer3 := NewConcreteObserver("Observer3")
	
	// subject.Attach(observer1)
	// subject.Attach(observer2)
	// subject.SetState("Hello, World!")
	
	// subject.Attach(observer3)
	// subject.SetState("Updated state")
	
	// subject.Detach(observer2)
	// subject.SetState("Final state")
	// fmt.Println()

	// // 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Stock Price Monitoring
	fmt.Println("Stock Price Monitoring System:")
	stockMarket := stocks.NewStockMarket()
	
	trader1 := stocks.NewStockTrader("Trader1")
	trader2 := stocks.NewStockTrader("Trader2")
	analyst1 := stocks.NewStockAnalyst("Analyst1")
	
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

	// // Weather Station
	// fmt.Println("Weather Station System:")
	// weatherStation := NewWeatherStation()
	
	// currentDisplay := NewCurrentConditionsDisplay("Display1")
	// statsDisplay := NewStatisticsDisplay("Stats1")
	// forecastDisplay := NewForecastDisplay("Forecast1")
	
	// weatherStation.Attach(currentDisplay)
	// weatherStation.Attach(statsDisplay)
	// weatherStation.Attach(forecastDisplay)
	
	// // Update weather measurements
	// weatherStation.SetMeasurements(25.0, 65.0, 1013.0)
	// weatherStation.SetMeasurements(32.0, 70.0, 1005.0)
	// weatherStation.SetMeasurements(-5.0, 80.0, 1020.0)
	// fmt.Println()

	// // Event System
	// fmt.Println("Event System:")
	// eventBus := NewEventBus()
	
	// loggingObserver := NewLoggingObserver("Logger1")
	// notificationObserver := NewNotificationObserver("Notifier1")
	
	// eventBus.Subscribe(loggingObserver)
	// eventBus.Subscribe(notificationObserver)
	
	// // Publish events
	// eventBus.Publish(&Event{
	// 	Type:      "user_login",
	// 	Data:      "User John logged in",
	// 	Timestamp: time.Now(),
	// 	Source:    "auth_service",
	// })
	
	// eventBus.Publish(&Event{
	// 	Type:      "error",
	// 	Data:      "Database connection failed",
	// 	Timestamp: time.Now(),
	// 	Source:    "database_service",
	// })
	
	// eventBus.Publish(&Event{
	// 	Type:      "warning",
	// 	Data:      "High memory usage detected",
	// 	Timestamp: time.Now(),
	// 	Source:    "monitoring_service",
	// })
	
	// eventBus.Publish(&Event{
	// 	Type:      "system_alert",
	// 	Data:      "Server maintenance scheduled",
	// 	Timestamp: time.Now(),
	// 	Source:    "admin_service",
	// })
	// fmt.Println()

	// fmt.Println("=== END OF DEMONSTRATION ===")
}
