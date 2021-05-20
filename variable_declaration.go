package main
import "fmt"

func main() {
	var x int      // declaring int variable, default value is 0
	fmt.Println("value of x after declaration but before assignment:",x)

	x = 10  // assigning value to variable
	fmt.Println("value of x after assignment:",x)

	var y int = 20  // direct initialization
	// var y = 20      //  same as above (without mentioning type, takes type based on value assigned)
	                   // depending on the value assigned, it will consider the type automatically
	var z = 40
	fmt.Println("y and z are :",y,z)
	fmt.Printf("%T", x)     // gives the datatype of the variable 
}