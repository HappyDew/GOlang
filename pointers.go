// package main
//
// import (
// 	"fmt"
// )
//
// func zero(xPtr *int) {
// 	*xPtr = 0
// }
// func main() {
// 	x := 5
// 	zero(&x)
// 	fmt.Println(x)
// }

// package main
//
// func square(x *float64) {
// 	*x = *x * *x
// }
// func main() {
// 	x := 1.5
// 	square(&x)
// }

// func one(xPtr *int) {
// 	*xPtr = 1
// }
// func main() {
// 	xPtr := new(int)
// 	one(xPtr)
// 	fmt.Println(*xPtr) // x is 1
// }

package main

import (
	"fmt"
)

func swap(x, y *int) {
	*x, *y = 1, 2
	*x, *y = *y, *x
}
func main() {
	x, y := new(int), new(int)
	swap(x, y)
	fmt.Println("x =", *x, "y =", *y)
}
