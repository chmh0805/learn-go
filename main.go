package main

import (
	"fmt"
	"strings"
)

/*
	We have to give the type of parameters.
	We can give it just one using comma, if the types of multiple parameter are same.
*/
func multiply(a, b int) int {
	return a * b;
}

/*
	We can give the type of parameter using '...' when It is array type.
*/
func repeatMe(words ...string) {
	fmt.Println(words);
}


/*
	Functions in Go can return multiple values.
*/
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name);
}

/*
	Naked Return
		-> If we set the name of return values, We can return them just writing "return".
*/
func lenAndLower(name string) (length int, lowercase string) {
	length = len(name);
	lowercase = strings.ToLower(name);
	return // length, lowercase
}

/*
	defer will be run after return or the function block ends.
*/
func deferFunc(name string) (upperName string) {
	defer fmt.Println("I'm done!"); // It will be run after retrun.
	upperName = strings.ToUpper(name);
	return
}

func superAdd(numbers ...int) int {
	total := 0;
	/*
		range {items} will return {index, item}
	*/
	for _, number := range numbers {
		total += number;
	}
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	return total;
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6);
	fmt.Println(result);
}