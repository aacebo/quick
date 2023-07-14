package value

import (
	"fmt"
)

type Float float64

func (self Float) Kind() Kind {
	return FLOAT
}

func (self Float) String() string {
	return fmt.Sprintf("%f", self)
}
