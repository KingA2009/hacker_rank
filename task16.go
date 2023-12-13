// Two friends Anna and Brian, are deciding how to split the bill at a dinner. Each will only pay for the items they consume. Brian gets the check and calculates Anna's portion. You must determine if his calculation is correct.
package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	sum := 0
	for i := range bill {
		if k == int32(i) {
			continue
		} else {
			sum += int(bill[i])
		}
	}
	if (sum / 2) == int(b) {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - int32(sum/2))
	}
}
