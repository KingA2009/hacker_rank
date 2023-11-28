// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.
package main

import "fmt"

func miniMaxSum(arr []int32) {
	var sumN []int64
	for i := range arr {
		var sum int64
		for j, number := range arr {
			if i == j {
				continue
			} else {
				sum += int64(number)
			}
		}
		sumN = append(sumN, sum)
	}
	var largestNumber, temp int64
	for _, element := range sumN {
		if int64(element) > temp {
			temp = int64(element)
			largestNumber = temp
		}
	}
	smallestNumber := sumN[0]

	for i := range sumN {
		if sumN[i] < smallestNumber {
			smallestNumber = sumN[i]
		}
	}
	fmt.Println(smallestNumber, largestNumber)
}
