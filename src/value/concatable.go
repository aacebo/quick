package value

type Concatable interface {
	Concat(Concatable) Concatable

	Comparable
}
