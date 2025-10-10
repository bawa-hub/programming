package database

import "fmt"

type Database struct {
	data map[string]interface{}
}

func NewDatabase() *Database {
	return &Database{
		data: make(map[string]interface{}),
	}
}

func (db *Database) Insert(key string, value interface{}) {
	db.data[key] = value
	fmt.Printf("Database: Inserted %s = %v\n", key, value)
}

func (db *Database) Update(key string, value interface{}) {
	if _, exists := db.data[key]; exists {
		db.data[key] = value
		fmt.Printf("Database: Updated %s = %v\n", key, value)
	}
}

func (db *Database) Delete(key string) {
	if _, exists := db.data[key]; exists {
		delete(db.data, key)
		fmt.Printf("Database: Deleted %s\n", key)
	}
}

func (db *Database) Get(key string) interface{} {
	return db.data[key]
}