package fruit

type FruitBasketIterator struct {
	basket *FruitsBasket
	index  int
}

func NewFruitBasketIterator(basket *FruitsBasket) *FruitBasketIterator {
	return &FruitBasketIterator{
		basket: basket,
		index:  0,
	}
}

func (it *FruitBasketIterator) HasNext() bool {
	return it.index < len(it.basket.fruits)
}

func (it *FruitBasketIterator) Next() *Fruit {
	if it.HasNext() {
		fruit := it.basket.fruits[it.index]
		it.index++
		return &fruit
	}
	return &Fruit{}
}
