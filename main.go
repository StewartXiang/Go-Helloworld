package main

import "fmt"

func main() {
	var i string
	i = "hello"
	i = i + "yes"
	const Pi float32 = 3.1415926
	var fslice = []int {1, 2, 3}
	slice := []byte {'a', 'b', 'c', 'd'}
	//var ar = [10]byte {'a', 'b', 'c'}
	fmt.Printf("Hello, world!\n")
	fmt.Printf(i+"\n")
	fmt.Printf("%f", Pi)
	fmt.Printf("%d, %s\n", fslice, slice)
	fmt.Println(cap(slice))
	var nslice = append(slice, 'e')
	fmt.Println(nslice, len(nslice))
	fmt.Println(cap(nslice))
Here:
	var nnslice = nslice[0:3]
	fmt.Println(len(nnslice))
	fmt.Println(cap(nnslice), nnslice[:4])
	goto Here
}
