package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[0] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dict:", b)
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	x := [...]int{2: 200, 4: 500}
	fmt.Println("idx:", x)
	var s []int
	s = append(s, 100) // Adds 100 to the slice
	s = append(s, 200) // Adds 200 to the slice
	s = append(s, 300) // Adds 300 to the slice
	fmt.Println("s", s)
	var mixedArray []interface{}

	mixedArray = append(mixedArray, 100)            // Add an int
	mixedArray = append(mixedArray, "Hello, world") // Add a string
	mixedArray = append(mixedArray, 3.14)           // Add a float
	mixedArray = append(mixedArray, true)           // Add a bool

	fmt.Println("Mixed Array:", mixedArray)
	for i, v := range mixedArray {
		fmt.Printf("Index %d: ", i)
		switch v := v.(type) { // Type assertion in a type switch
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case float64:
			fmt.Println("Float:", v)
		case bool:
			fmt.Println("Boolean:", v)
		default:
			fmt.Println("Unknown type")
		}
	}
	for i, v := range mixedArray {
		fmt.Println("index", i, " there is the val", v)
	}
	for _, v := range mixedArray {
		fmt.Println(" there is the val", v)
	}
	for i := range mixedArray {
		fmt.Println(" there is the val", i)
	}
}
