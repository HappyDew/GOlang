// package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// func distance(x1, y1, x2, y2 float64) float64 {
// 	a := x2 - x1
// 	b := y2 - y1
// 	return math.Sqrt(a*a + b*b)
// }
// func rectangleArea(x1, y1, x2, y2 float64) float64 {
// 	l := distance(x1, y1, x1, y2)
// 	w := distance(x1, y1, x2, y1)
// 	return l * w
// }
// func circleArea(x, y, r float64) float64 {
// 	return math.Pi * r * r
// }
// func main() {
// 	var rx1, ry1 float64 = 0, 0
// 	var rx2, ry2 float64 = 10, 10
// 	var cx, cy, cr float64 = 0, 0, 5
//
// 	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
// 	fmt.Println(circleArea(cx, cy, cr))
// }

// package main
//
// import "fmt"
//
// func main() {
// 	x := 1
// 	x1 := 2
// 	fmt.Println(x, x1)
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

// func main() {
// 	go f(0)
// 	var input string
// 	fmt.Scanln(&input)
// }
