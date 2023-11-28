// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with 6 places after the decimal.
//
// Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  are acceptable.
package main

import "fmt"

func plusMinus(arr []int32) {
	var plusN, minusN, zeroN int
	for _, i := range arr {
		if i > 0 {
			plusN++
		} else if i < 0 {
			minusN++
		} else {
			zeroN++
		}
	}
	fmt.Println(fmt.Sprintf("%.6f", float64(plusN)/float64(len(arr))))
	fmt.Println(fmt.Sprintf("%.6f", float64(minusN)/float64(len(arr))))
	fmt.Println(fmt.Sprintf("%.6f", float64(zeroN)/float64(len(arr))))
}
