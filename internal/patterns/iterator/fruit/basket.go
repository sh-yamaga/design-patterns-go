package fruit

type Basket struct {
	fruits []Fruit
}

func NewBasket(fs []Fruit) *Basket {
	return &Basket{
		fruits: fs,
	}
}

func (fb Basket) At(i int) *Fruit {
	return &fb.fruits[i]
}

func (fb *Basket) Append(f Fruit) {
	fb.fruits = append(fb.fruits, f)
}

func (fb Basket) Length() int {
	return len(fb.fruits)
}

func (fb Basket) Iterator() *basketIterator {
	return newBasketIterator(&fb)
}
