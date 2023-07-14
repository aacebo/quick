package value

type Comparable interface {
	comparable
	Value

	Byte | Bool | Int | Float | String
}
