// package main
//
// func average(xs []float64) float64 {
// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	return total / float64(len(xs))
// }

// package main
//
// import "fmt"
//
// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}
// 	fmt.Println(average(xs))
// }

// package main
//
// import "fmt"
//
// func main() {
// 	fmt.Println(f1())
// }
// func f1() int {
// 	return f2()
// }
// func f2() (r int) {
// 	r = 1
// 	return
// }

// package main
//
// func f() (int, int) {
// 	return 5, 6
// }
// func main() {
// 	x, y := f() //почемуто не работает
// }

// package main
//
// import "fmt"

// func add(args ...int) int {
// 	total := 0
// 	for _, v := range args {
// 		total += v
// 	}
// 	return total
// }
// func main() {
// 	fmt.Println(add(1, 2, 3, 3))
// }

// package main
//
// import "fmt"
//
// func main() {
// 	add := func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println(add(1, 1))
// }

package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
