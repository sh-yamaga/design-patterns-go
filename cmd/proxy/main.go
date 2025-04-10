package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/proxy"
)

func main() {
	// generate proxyDatabase
	proxyDB := proxy.NewProxyDatabase("Proxy Sample DB")

	// 1st Query
	fmt.Println("=== 1st Query ===")
	result, _ := proxyDB.Query("SELECT * FROM users;")
	fmt.Println("Result:", result)
	// Output:
	// === 1st Query ===
	// RealDatabase[Proxy Sample DB] executing query: SELECT * FROM users;
	// Querying RealDatabase...
	// Result: RealDatabaseResult

	// 2nd Query
	fmt.Println("=== 2nd Query ===")
	result, _ = proxyDB.Query("SELECT * FROM users;")
	fmt.Println("Result:", result)
	// Output:
	// === 2nd Query ===
	// Proxy: using cache.
	// Result: RealDatabaseResult

	// 3rd Query
	fmt.Println("=== 3rd Query ===")
	result, _ = proxyDB.Query("SELECT * FROM users;")
	fmt.Println("Result:", result)
	// Output:
	// === 3rd Query ===
	// Proxy: using cache.
	// Result: RealDatabaseResult
}
