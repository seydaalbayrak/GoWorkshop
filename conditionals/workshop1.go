package conditionals
import "fmt"
func Demo1()  {
	var num1 int = 1
	var num2 int = 9
	var num3 int = 15

	var biggest int=num1

	if num2>biggest {
		biggest=num2
	}
	if num3>biggest{
		biggest=num3
	}
	fmt.Printf("Biggest number: %v",biggest)
	
}