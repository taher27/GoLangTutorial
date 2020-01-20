package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		temp := f1 + f2
		f1 = f2
		f2 = temp
		return f2
	}
}

func main() {
	f := fibonacci()
	fmt.Print("0 1 ")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
}
