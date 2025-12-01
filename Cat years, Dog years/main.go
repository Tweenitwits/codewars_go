/*
I have a cat and a dog.

I got them at the same time as kitten/puppy.
That was humanYears years ago.

Return their respective ages now as [humanYears,catYears,dogYears]

NOTES:

humanYears >= 1
humanYears are whole numbers only
Cat Years
15 cat years for first year
+9 cat years for second year
+4 cat years for each year after that
Dog Years
15 dog years for first year
+9 dog years for second year
+5 dog years for each year after that
*/

package kata

func CalculateYears(years int) (result [3]int) {
	var catYears, dogYears int
	if years == 1 {
		catYears = 15
		dogYears = 15
	} else if years > 1 && years < 3 {
		catYears = 24
		dogYears = 24
	} else {
		catYears = 24 + (years-2)*4
		dogYears = 24 + (years-2)*5
	}
	result[0] = years
	result[1] = catYears
	result[2] = dogYears
	return result
}
