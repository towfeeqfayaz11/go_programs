package main
import "fmt"

func main() {
	a := 10
	fmt.Printf("hello, type is: %T and value is %v\n", a, a)       // %T is to print type and %v is to print value
	var s string = fmt.Sprintf("hello, type is: %T and value is %v", a, a)   // same as printf, except rather than print the output to console, it will store it in variable

	fmt.Println("printing varibale s now:")
	fmt.Println(s)

	fmt.Printf("the boolean is : %t\n", true)    // golang used "true" or false
}
