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

func main() {
	fmt.Println(multiply(3, 4));

	/*
		We can ignore the value of return values with using '_'.
	*/
	totalLength, _ := lenAndUpper("hyuk");
	fmt.Println(totalLength);

	repeatMe("hyuk", "sun", "123", "456");
}