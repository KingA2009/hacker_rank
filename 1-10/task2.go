// Given an array of integers, find the sum of its elements.
//
// For example, if the array a = [1,2,3] ,1+2+3=6 , so return 6.
package main

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	var sum int32
	for i := range ar {
		sum = sum + ar[i]
	}
	return sum
}
