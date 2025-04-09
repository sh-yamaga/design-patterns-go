package proxy

import (
	"fmt"
	"time"
)

type Database interface {
	Query(query string) (string, error)
}

// RealDatabase represents a concrete implementation of the Database interface
type RealDatabase struct {
	Name string
}

// Query executes the given query and returns a simulated result
func (db *RealDatabase) Query(query string) (string, error) {
	fmt.Printf("RealDatabase[%s] executing query: %s\n", db.Name, query)

	// Simulate a heavy task with a sleep
	fmt.Print("Querying RealDatabase")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println()

	// Simulated result
	result := "RealDatabaseResult"

	return result, nil
}
