package fruit

type BasketIterator struct {
	basket *Basket
	index  int
}

func newBasketIterator(basket *Basket) *BasketIterator {
	return &BasketIterator{
		basket: basket,
		index:  0,
	}
}

func (ittr *BasketIterator) HasNext() bool {
	return ittr.index < len(ittr.basket.fruits)
}

func (ittr *BasketIterator) Next() *Fruit {
	if ittr.HasNext() {
		fruit := ittr.basket.fruits[ittr.index]
		ittr.index++
		return &fruit
	}
	return &Fruit{}
}
