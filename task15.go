// Given an array of bird sightings where every element represents a bird type id, determine the id of the most frequently sighted type. If more than 1 type has been spotted that maximum amount, return the smallest of their ids.
package main

func migratoryBirds(arr []int32) int32 {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[int(num)] = freq[int(num)] + 1
	}
	var largestNumber, temp int64
	for _, element := range freq {
		if int64(element) > temp {
			temp = int64(element)
			largestNumber = temp
		}

	}
	var result, sum int32
	for i, element := range freq {
		if largestNumber == int64(element) {
			result = int32(i)
			sum++
		}
	}
	if sum > 1 {
		for i, element := range freq {
			if largestNumber == int64(element) && result > int32(i) {
				result = int32(i)
			}
		}
	}
	return result
}
