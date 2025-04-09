package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/proxy"
)

func main() {
	// generate RealDatabase
	realDB := &proxy.RealDatabase{
		Name: "Proxy Sample Database",
	}
	// generate proxyDatabase
	proxyDB := proxy.NewProxyDatabase(realDB)

	// 1st Query
	fmt.Println("=== 1st Query ===")
	result, _ := proxyDB.Query("SELECT * FROM users;")
	fmt.Println("Result:", result)

	// 2nd Query
	fmt.Println("=== 2nd Query ===")
	result, _ = proxyDB.Query("SELECT * FROM users;")
	fmt.Println("Result:", result)

	// 3rd Query
	fmt.Println("=== 3rd Query ===")
	result, _ = proxyDB.Query("SELECT * FROM users;")
	fmt.Println("Result:", result)
}
