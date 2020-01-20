package main

func swap(p *int, q *int) {
	var temp int
	temp = *p
	*p = *q
	*q = temp
}

//*************************Main function******************************
// func main() {

// 	var a int
// 	var b int
// 	a = 10
// 	b = 20

// 	fmt.Println("Before swap")
// 	fmt.Printf("A: %d, B: %d \n", a, b)
// 	swap(&a, &b)
// 	fmt.Println("After swap")
// 	fmt.Printf("A: %d, B: %d \n", a, b)
// }
