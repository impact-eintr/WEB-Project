package b

import (
	"fmt"
	a "test/a"
)

func B() {
	fmt.Println("This is B func")
	a.A()
}
