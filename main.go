package helloworld

import "fmt"

func main() {
//	var i string
//	i = "hello"
//	i = i + "yes"
//	const Pi float32 = 3.1415926
//	var fslice = []int {1, 2, 3}
//	slice := []byte {'a', 'b', 'c', 'd'}
//	//var ar = [10]byte {'a', 'b', 'c'}
//	fmt.Printf("Hello, world!\n")
//	fmt.Printf(i+"\n")
//	fmt.Printf("%f", Pi)
//	fmt.Printf("%d, %s\n", fslice, slice)
//	fmt.Println(cap(slice))
//	var nslice = append(slice, 'e')
//	fmt.Println(nslice, len(nslice))
//	fmt.Println(cap(nslice))
//Here:
//	var nnslice = nslice[0:3]
//	fmt.Println(len(nnslice))
//	fmt.Println(cap(nnslice), nnslice[:4])
//	goto Here


	var P person
	P.name = "Xiang"
	P.age = 20
	//fmt.Printf("He is %s", P.name)
	var names = [6]string {
		"aaaa", "bbbb", "cccc", "dddd", "eeee", "ffff",
	}
	var ages = [6]int {22, 23, 24, 25, 26, 27}
	var my_persons [][2]person
	for i, n := range names{
		if i%2 == 0{
			my_persons = append(
				my_persons, [2]person {{n, ages[i]}, })
		}else{
			my_persons[len(my_persons)-1][1] = person{n, ages[i]}
		}
	}
	fmt.Println(my_persons)
	f := filter(my_persons, Older)
	fmt.Println(f)
}

type age_compare func(p1, p2 person) person
type person struct {
	name string
	age int
}

func Older(p1, p2 person) person {
	if p1.age>p2.age {
		return p1
	}
	return p2
}

func filter(persons [][2]person, f age_compare) []person{
	var result []person
	for _, group := range persons{
		result = append(result, f(group[0], group[1]))
	}
	return result
}
