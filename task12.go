//	Submissions	Leaderboard	Discussions	Editorial
//
// HackerLand University has the following grading policy:
//
// Every student receives a  in the inclusive range from  to .
// Any  less than  is a failing grade.
// Sam is a professor at the university and likes to round each student's  according to these rules:
//
// If the difference between the  and the next multiple of  is less than , round  up to the next multiple of .
// If the value of  is less than , no rounding occurs as the result will still be a failing
package main

func gradingStudents(grades []int32) []int32 {
	var result []int32
	for _, number := range grades {
		if number < 38 {
			result = append(result, number)
		} else if number%5 != 0 {
			n := number / 5
			h := (n + 1) * 5
			if h-number < 3 {
				result = append(result, h)
			} else {
				result = append(result, number)
			}
		} else {
			result = append(result, number)
		}
	}
	return result
}
