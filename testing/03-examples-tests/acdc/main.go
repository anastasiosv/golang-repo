//Package acdc provides func to sum the ints
package acdc

//Sum is the function that provides the outcome of the sum process
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v

	}
	return s
}
