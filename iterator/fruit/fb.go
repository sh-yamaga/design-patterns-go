package fruit

type FruitsBasket struct {
	fruits []Fruit
}

func NewFruitBasket(fs []Fruit) *FruitsBasket {
	return &FruitsBasket{
		fruits: fs,
	}
}

func (fb FruitsBasket) At(i int) *Fruit {
	return &fb.fruits[i]
}

func (fb *FruitsBasket) Append(f Fruit) {
	fb.fruits = append(fb.fruits, f)
}

func (fb FruitsBasket) Length() int {
	return len(fb.fruits)
}

func (fb FruitsBasket) Iterator() *FruitBasketIterator {
	return &FruitBasketIterator{
		basket: &fb,
		index:  0,
	}
}
