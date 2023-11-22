//In this challenge, you are required to calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large.

package main

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, a := range ar {
		sum += a
	}
	return sum

}
