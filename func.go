// package main
//
// import "fmt"

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

// package main
//
// import "fmt"

// func main() {
// 	x := 0
// 	increment := func() int {
// 		x++
// 		return x
// 	}
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

// package main
//
// import "fmt"
//
// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// func makeEvenGenerator() func() uint { //замыкания
// 	i := uint(0)
// 	return func() (ret uint) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }
// func main() {
// 	nextEven := makeEvenGenerator()
// 	fmt.Println(nextEven()) //0
// 	fmt.Println(nextEven()) //2
// 	fmt.Println(nextEven()) //4
// }

//Рекурсия
// package main
// import "fmt"
//
// func factorial(x uint) uint{
//   if x == 0 {
//     return 1
//   }
//   return x * factorial(x-1)
// }

// package main
//
// import "fmt"
//
// func first() {
// 	fmt.Println("1st")
// }
// func second() {
// 	fmt.Println("2st")
// }
// func main() {
// 	defer second() //это ключевое слово  типо "потом"
// 	first()
// }

// package main
//
// import "fmt"
//
// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}()
// 	panic("PANIC")
// }
package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 6333, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	max := x[0]
	for _, value := range x {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}

// func main(){
//   x := []float64{2, 3, 4, 5, 66, 77}
//   max := 99
//   for _, i := range x[]{
//     if max < x[i] {
//       max = x[i]
//     }
//     fmt.Println(max)
//   }
// }
