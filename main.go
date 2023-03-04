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

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 { // can make local variable in conditional loop
	// 	return false;
	// }
	switch koreanAge := age + 2; koreanAge { // can also make local variable in switch ...case loop
	case 10:
		return false
	case 18:
		return true
	}

	return false;
}

func main() {
	names := [5]string{"hyuk", "sun", "mun"}; // Array
	// can make array like this -> [...]string{"a", "b", "c"};
	names[3] = "aaa";
	names[4] = "bbb";
	// names[5] = "asdasd"; // out of bounds.

	ages := []int{18, 19, 20}; // Slice
	// ages[3] = 21; // cannot add item like this.
	ages = append(ages, 21); // append doesn't modify origin slice, just return modified slice.

	fmt.Println(names);
	fmt.Println(ages);
}