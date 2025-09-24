package main

import (
	"fmt"
)

func main() {

	// 1. ENCAPSULATION
	fmt.Println("1. ENCAPSULATION:")
	account := NewBankAccount("12345", "John Doe", 1000.0)
	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
	
	err := account.Deposit(500.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After deposit: $%.2f\n", account.GetBalance())
	}

	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After withdrawal: $%.2f\n", account.GetBalance())
	}
	fmt.Println()

	// 2. INHERITANCE
	fmt.Println("2. INHERITANCE:")
	car := &Car{
		Vehicle: Vehicle{Brand: "Toyota", Model: "Camry", Year: 2023},
		Doors:   4,
		Engine:  "V6",
	}
	car.Start()
	car.OpenTrunk()

	motorcycle := &Motorcycle{
		Vehicle:       Vehicle{Brand: "Honda", Model: "CBR600", Year: 2023},
		HasWindshield: true,
	}
	motorcycle.Start()
	motorcycle.Wheelie()
	fmt.Println()

	// 3. POLYMORPHISM
	fmt.Println("3. POLYMORPHISM:")
	vehicles := []Startable{car, motorcycle}
	for _, vehicle := range vehicles {
		StartVehicle(vehicle)
	}
	fmt.Println()

	// 4. ABSTRACTION
	fmt.Println("4. ABSTRACTION:")
	shapes := []Shape{
		&Rectangle{Width: 5, Height: 3},
		&Circle{Radius: 4},
	}
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
	fmt.Println()

	// 5. SOLID PRINCIPLES
	fmt.Println("5. SOLID PRINCIPLES:")
	
	// SRP - Separate concerns
	user := &User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	emailService := &EmailService{}
	fmt.Printf("User: %s\n", user.GetDisplayName())
	emailService.SendEmail(user.Email, "Welcome", "Welcome to our service!")
	fmt.Println()

	// OCP - Open for extension
	creditCardProcessor := &CreditCardProcessor{CardNumber: "1234-5678-9012-3456"}
	paymentService := &PaymentService{processor: creditCardProcessor}
	paymentService.ProcessPayment(100.0)
	fmt.Println()

	// LSP - Substitutability
	birds := []Bird{&Sparrow{}, &Penguin{}}
	for _, bird := range birds {
		MakeBirdFly(bird)
	}
	fmt.Println()

	// ISP - Interface segregation
	human := &Human{Name: "Bob"}
	robot := &Robot{Model: "T-800"}
	
	workers := []Worker{human, robot}
	for _, worker := range workers {
		worker.Work()
	}
	fmt.Println()

	// DIP - Dependency inversion
	mysqlDB := &MySQLDatabase{}
	userService := &UserService{db: mysqlDB}
	userService.CreateUser("Charlie")
	userService.GetUser("1")
	fmt.Println()

	// 6. COMPOSITION VS INHERITANCE
	fmt.Println("6. COMPOSITION VS INHERITANCE:")
	modernCar := &ModernCar{
		Brand:  "BMW",
		Model:  "X5",
		Engine: Engine{Type: "V8", Horsepower: 400},
	}
	modernCar.Start()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
