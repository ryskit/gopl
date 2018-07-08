package main

import "fmt"

func main() {
	z := make([]int, 5, 10)
	fmt.Printf("len = %d\n", cap(z))
	fmt.Printf("cap = %d\n", len(z))
	z[0] = 5
	fmt.Println(z[0])
	fmt.Println(z[1])
	fmt.Println(z[2])
	fmt.Println(z[3])
	fmt.Println(z[4])
	x := make([]int, 10, 20)
	copy(x, z)
	fmt.Println(x)
}
