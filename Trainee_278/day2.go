package main

func getAverage(arr []int, size int) float32 {
	sum := 0
	for i := 0; i < size; i++ {
		sum += arr[i]
	}

	return float32(sum / size)
}

//*************************Main function******************************

// func main() {
// 	var greet = "Hello World!"
// 	fmt.Println(greet)

// 	for i := 0; i < len(greet); i++ {
// 		fmt.Printf("%c = %x,\n", greet[i], greet[i])
// 	}
// 	Length of string
// 	fmt.Println("Length of %s is %d", greet, len(greet))

// 	Join of 2 strings
// 	weekdays := []string{"Monday", "Tuesday", "Wednessday"}
// 	fmt.Println(strings.Join(weekdays, " "))

// 	var bal = []int{2, 4, 5, 6, 7, 8, 9, 10}
// 	var avg float32

// 	avg = getAverage(bal, len(bal))
// 	fmt.Printf("The average is %f.\n", avg)

// 	var fptr *float32
// 	fptr = &avg

// 	fmt.Printf("The address is %x.\n", fptr)

// 	var val int
// 	var ptr *int
// 	var pptr **int

// 	val = 1000
// 	ptr = &val
// 	pptr = &ptr

// 	fmt.Printf("value of val: %d\n", val)
// 	fmt.Printf("value of ptr1: %x\n", &ptr)
// 	fmt.Printf("value of ptr2: %x\n", ptr)
// 	fmt.Printf("value of pptr: %x\n", *pptr)

// }
