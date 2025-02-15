package fruit

type basketIterator struct {
	basket *Basket
	index  int
}

func newBasketIterator(basket *Basket) *basketIterator {
	return &basketIterator{
		basket: basket,
		index:  0,
	}
}

func (ittr *basketIterator) HasNext() bool {
	return ittr.index < len(ittr.basket.fruits)
}

func (ittr *basketIterator) Next() *Fruit {
	if ittr.HasNext() {
		fruit := ittr.basket.fruits[ittr.index]
		ittr.index++
		return &fruit
	}
	return &Fruit{}
}
