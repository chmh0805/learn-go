package main

import "fmt"

func main() {
	const name string = "hyuk";

	/*
		var_name := var_value
			-> Go will infer the type of variable.
			-> The type of value cannot be changed after been inferred.
			-> This grammer cannot be used outside of function block.
	*/
	age := 27; // var age int16 = 27;

	fmt.Printf("I'm %s. and %d years old.\n", name, age);
}