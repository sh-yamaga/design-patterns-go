package element

type Element interface {
	Accept(Visitor)
}
