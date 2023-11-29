// Mark and Jane are very happy after having their first child. Their son loves toys, so Mark wants to buy some. There are a number of different toys lying in front of him, tagged with their prices. Mark has only a certain amount to spend, and he wants to maximize the number of toys he buys with this money. Given a list of toy prices and an amount to spend, determine the maximum number of gifts he can buy.
package main

import "sort"

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices[:], func(i, j int) bool {
		return prices[:][i] < prices[:][j]
	})
	var sumPrices, sum int32
	for _, price := range prices {
		if k > sumPrices {
			sumPrices = sumPrices + price
			sum++

		}
	}
	return sum - 1
}
