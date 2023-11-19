// Complete the function solveMeFirst to compute the sum of two integers.
// Complete the solveMeFirst function in the editor below.
//
// solveMeFirst has the following parameters:
//
// int a: the first value
// int b: the second value
// Returns
// - int: the sum of a and b
package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	return (a + b)
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
