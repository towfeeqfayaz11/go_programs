package main
import "fmt"

// function without argument and  without return type
func main() {
	fmt.Println("starting main function")
	displayMessage("hi towfeeq")
	displayValue(20)
	fmt.Println("ending main function")


}


// function with argument but no return type

func displayMessage(s string) {
	fmt.Println("inside displayMessage function")
	fmt.Println("Message is:",s)
	fmt.Printf("%s\n",s)      // c style print function
}

// function with argument but no return type
func displayValue(i int) {
	fmt.Println("inside displayValue function")
	fmt.Println("value is:",i)
	fmt.Printf("%d\n",i)      // c style print function
}