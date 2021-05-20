package main
import "fmt"

// function without argument and without return type
func main() {
	var result int
	result = add_numbers(10,20)
	fmt.Println("sum of no.s is :",result)

	// using c-style print fucntion
	fmt.Printf("result is : %d \n", result)
}


// function with argument and with return type
func add_numbers(x int, y int) int {        // here last int is return type
	var z int
	z = x + y
	return z
}