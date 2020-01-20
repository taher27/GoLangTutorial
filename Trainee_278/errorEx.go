package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt struct {
	val float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",e.val)
}

func Sqrt(x float64) (float64, error) {
	if x<0 {
		return x, &ErrNegativeSqrt { x }
	}
	return math.Sqrt(x),nil
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
