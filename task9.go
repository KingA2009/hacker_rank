// You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.
package main

func birthdayCakeCandles(candles []int32) int32 {
	var largestNumber, temp int64
	for _, element := range candles {
		if int64(element) > temp {
			temp = int64(element)
			largestNumber = temp
		}
	}
	var sum int32
	for _, i := range candles {
		if largestNumber == int64(i) {
			sum++
		}
	}
	return sum
}
