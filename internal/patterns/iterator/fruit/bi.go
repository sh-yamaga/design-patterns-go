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

func (bi *basketIterator) HasNext() bool {
	return bi.index < bi.basket.Length()
}

func (bi *basketIterator) Next() *Fruit {
	if bi.HasNext() {
		fruit := bi.basket.At(bi.index)
		bi.index++
		return fruit
	}
	return &Fruit{}
}
