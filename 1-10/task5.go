// Complete the  function in the editor below.
//
// diagonalDifference takes the following parameter:
//
// int arr[n][m]: an array of integers
// Return
//
// int: the absolute diagonal difference
package main

func diagonalDifference(arr [][]int32) int32 {
	var suma, sumb int32
	for i := range arr {
		suma += arr[i][i]
		sumb += arr[i][len(arr)-i-1]
	}
	var result int32
	if suma > sumb {
		result = suma - sumb
	} else if suma < sumb {
		result = sumb - suma
	} else {
		result = 0
	}
	return result
}
