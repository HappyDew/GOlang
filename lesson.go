// package main
// import "fmt"
//
// func main(){
//   var x [5]int
//   x[4] = 100
//   fmt.Println(x)
// }
//
// package main
// import "fmt"
//
// func main(){
//   x := [5]float64{ 98, 93, 77, 82, 83 }
//
//   var total float64 = 0
//   for _, value := range x {
//     total += value
//   }
//   fmt.Println(total / float64(len(x)))
// }

// package main
// import "fmt"
//
// func main(){
//   slice1 := []int{1,2,3}
//   slice2 := append(slice1, 4, 5)
//   fmt.Println(slice1, slice2)
// }

// package main
// import "fmt"
//
// func main(){
//   x := make(map[string]int)
//   x["key"] = 10
//   fmt.Println(x["key"])
// }

// package main
// import "fmt"
//
// func main(){
//   elements := make(map[string]string)
//   elements["H"] = "Hydrogen" // "H": "Hydrogen",
//   elements["He"] = "Helium"
//   elements["Li"] = "Lithium"
//   elements["Be"] = "Beryllium"
//   elements["B"] = "Boron"
//   elements["C"] = "Carbon"
//   elements["N"] = "Nitrogen"
//   elements["O"] = "Oxygen"
//   elements["F"] = "Fluorine"
//   elements["Ne"] = "Neon"
//
//   name, ok := elements["Un"]
//   fmt.Println(name, ok)
// }

package main

import (
	"fmt"
	"sort"
) //добавляется пакет

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	// min := x[0]
	// for i := 0; i < len(x); i++ {
	//   if min > x[i] {
	//     min = x[i]
	//   }
	// }
	sort.Ints(x)
	fmt.Println( /*min*/ x[0])
}
