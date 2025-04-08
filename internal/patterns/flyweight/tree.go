package flyweight

import "fmt"

// TreeType holds Tree's features
type TreeType struct {
	Name    string
	Color   string
	Texture string
}

func (tt *TreeType) Draw(x, y int) {
	fmt.Printf("Name: %s, Color: %s, Texture: %s\nCoordinate: (%d, %d)\n", tt.Name, tt.Color, tt.Texture, x, y)
}

// Tree holds internal(Type) and external(x, y) features
type Tree struct {
	X, Y int
	Type *TreeType
}

func (t *Tree) Draw() {
	t.Type.Draw(t.X, t.Y)
}
