package main         // Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
import "fmt"         // one of standard package which helps us to print to console
                     // fmt package, which contains functions for formatting text, including printing to the console. 
                     // This package is one of the standard library packages you got when you installed Go.


// function with no argument and no return type
func main() {        // Implement a main function to print a message to the console. A main function executes by default when you run the main package.
	fmt.Println("hello golang")
	testfunction()    // calling a custom function
	fmt.Println("finishing code")
}

// function with no argument and no return type
func testfunction() {
	fmt.Print("executing test funcntion")
	fmt.Println("bye bye from test func")
}