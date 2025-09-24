package main

import "fmt"

// 5. DEPENDENCY INVERSION PRINCIPLE (DIP)
// High-level modules should not depend on low-level modules

// Database interface (abstraction)
type Database interface {
	Save(data string) error
	Get(id string) (string, error)
}

// MySQLDatabase implements Database
type MySQLDatabase struct{}

func (mdb *MySQLDatabase) Save(data string) error {
	fmt.Printf("Saving to MySQL: %s\n", data)
	return nil
}

func (mdb *MySQLDatabase) Get(id string) (string, error) {
	fmt.Printf("Getting from MySQL: %s\n", id)
	return "data", nil
}

// UserService depends on Database interface, not concrete implementation
type UserService struct {
	db Database
}

func (us *UserService) CreateUser(name string) error {
	return us.db.Save(name)
}

func (us *UserService) GetUser(id string) (string, error) {
	return us.db.Get(id)
}

func main() {
	mysqlDB := &MySQLDatabase{}
	userService := &UserService{db: mysqlDB}
	userService.CreateUser("Charlie")
	userService.GetUser("1")
	fmt.Println()
}
