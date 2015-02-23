/*z = z-((z**z-x)2z)/*/

package main

import "fmt"

func calc(x float64, z float64) float64 {
	z_square := z * z
	return z - ((z_square - x) / (2 * z))
}

func main() {
	/* Implements Newton's method of finding square root of a number*/
	delta := 0.00000005
	diff := 1.0
	z := 10.0
	old_z := z
	i := 1
	for diff > delta {
		z = calc(2.0, old_z)
		diff = old_z - z
		old_z = z
		i += 1
	}
	fmt.Println(z)
	fmt.Println(i)
}
