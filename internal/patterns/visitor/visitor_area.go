package visitor

import (
	"fmt"
	"math"
)

type AreaVisitor struct{}

func (av *AreaVisitor) VisitCircle(c *Circle) {
	area := math.Pi * c.Radius * c.Radius
	fmt.Printf("AreaVisitor: Area of Circle = %.2f\n", area)
}

func (av *AreaVisitor) VisitRectangle(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("AreaVisitor: Area of Rectangle = %.2f\n", area)
}
