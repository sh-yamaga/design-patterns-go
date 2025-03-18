package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/builder"
)

func main() {
	hb := builder.NewHtmlBuilder()
	html := hb.SetTitle("Builder Pattern").
		AddParagraph("This is a sample program that demonstrates the Builder Pattern, a type of GoF Design Pattern").
		AddList([]string{"item1", "item2", "item3"}).
		Build()

	fmt.Println(html)
	// Output:
	// <html>
	// <head>
	// <title>Builder Pattern</title>
	// </head>
	// <body>
	// <p>This is a sample program that demonstrates the Builder Pattern, a type of GoF Design Pattern</p>
	// <ul>
	// <li>item1</li>
	// <li>item2</li>
	// <li>item3</li>
	// </ul>
	// </body>
	// </html>

	sb := builder.NewSQLBuilder()
	query := sb.Select("id", "name", "age").
		From("users").
		Where("age > 18").
		Build()

	fmt.Println(query)
	// Output:
	// SELECT id, name, age FROM users WHERE age > 18;
}
