package fruit

type Fruit struct {
	name  string
	color string
}

func NewFruit(name, color string) *Fruit {
	return &Fruit{
		name:  name,
		color: color,
	}
}

func (f Fruit) Name() string {
	return f.name
}

func (f Fruit) Color() string {
	return f.color
}
