// Number line jumps
package main

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 > x2 {
		if v2 > v1 && (x1-x2)%(v2-v1) == 0 {
			return "YES"
		} else {
			return "NO"
		}
	} else if x1 < x2 {
		if v1 > v2 && (x2-x1)%(v1-v2) == 0 {
			return "YES"
		} else {
			return "NO"
		}
	} else {
		return "YES"
	}
}
