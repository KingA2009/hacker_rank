// Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.
//
// Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
package main

import "time"

func timeConversion(s string) string {
	t, _ := time.Parse("03:04:05PM", s)
	return (t.Format("15:04:05"))
}
