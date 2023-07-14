package value

import (
	"fmt"
)

type Int int64

func (self Int) Kind() Kind {
	return INT
}

func (self Int) String() string {
	return fmt.Sprintf("%d", self)
}
