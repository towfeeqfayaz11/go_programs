package main
import "fmt"

func main() {
	fmt.Printf("the value is : %f \n", 2387.1678438123457435)         // by default prints upto 6 decimal places
	fmt.Printf("the value is : %g \n", 2387.1678438123457435)         // %g is used to print large decimal float (upto 12 decimals)
	// // to modify the the floats, we cna use width and precision
	/*

	Width and Precision

	%f    --> (default width, default precision)    // width means total length of the result, while precision means no. of digits after decimal
	%9f   --> (width 9, default precision)
	%.2f  --> (default width, precision 2)
	%9.2f --> (width 9, precision 2)
	%9.f  --> (width 9, precision 0)   // means print whole no., with zero decimal  (same as converting float to int)

	*/

	fmt.Printf("the value is : %6.8f \n", 2387.1678438123457435)


	// using padding along with width and precision
	// default padding is done using spaces

	fmt.Printf("The stirng is:%9s \n", "tim")              // %s is for string and %q is to print string along with double quotes
	                                                       // string is by default right jsutified

	fmt.Printf("The stirng is:%-9sis a tough guy\n", "tim")                 // print string using padding with left justified (- for left justification)
	fmt.Printf("The stirng is:%09s*is*a*tough*guy \n", "tim")        // do passing with 0 rather than default space

}
