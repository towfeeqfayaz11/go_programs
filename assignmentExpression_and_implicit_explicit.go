package main
import "fmt"


func main() {
	// explicit type for a variable
	var x int = 10       // explicit type dclaration

	var y int            // explicit type declaration
	y = 20

	var z = 30       // implicit type declaration // this variable can be reassigned any other value later (only int because implicit type as per assigned value is int)

	var w = 33.47   // implicit type declaration  // this variable can be reassigned any other value later (only float because implicit type as per assigned value is float)

	v := 33     // := is called expression assignment operator or vulrus operator
	            // it is same as implicit type declaration except that we dont need to write "var" key word before variable
    v = 54      // reassigning value
    p := "john"
    p = "mike"   // reassigning value 


    // note below are differnt ways of declaring and assigning variable

    // 1) var a int
    //    a = 10

    // 2) var a int = 10

    // 3) var a = 10

    // 4) a := 10

    //here 1) , 2) , 3), 4) all mean same thing


    fmt.Println(x,y,z,w,v,p)
}