package visitor

type Element interface {
	Accept(Visitor)
}
