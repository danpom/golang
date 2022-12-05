/*leap is used to determine if a given int is leap year*/
package leap

// IsLeapYear takes in a year(int) and determines if the year is a leap year
func IsLeapYear(year int) bool {

	// 	on every year that is evenly divisible by 4
	//   except every year that is evenly divisible by 100
	//     unless the year is also evenly divisible by 400
	switch {
	case year%4 == 0 && year%400 == 0 && year%100 == 0:
		return true
	case year%4 == 0 && year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
