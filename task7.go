// Staircase detail
//
// This is a staircase of size n = 4
//
//	 #
//	##
//
// ###
// ####
package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		var k string
		k = strings.Repeat(" ", int(n)-i-1)
		for j := i; j >= 0; j-- {
			k += "#"
		}
		fmt.Println(k)
	}

}
