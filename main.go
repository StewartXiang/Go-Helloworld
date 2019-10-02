package main

import (
	"math/rand"
	"fmt"
	"time"
)
func init(){
	rand.Seed(time.Now().UnixNano())
}

func main() {

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
				my_persons,
				[2]person {{n, ages[i], get_color(rand.Intn(6))}})
		}else{
			my_persons[len(my_persons)-1][1] = person{n, ages[i], get_color(rand.Intn(6))}
		}
	}
	fmt.Println(my_persons)
	f := filter(my_persons, Older)
	fmt.Println(f)
}


const(
	WHITE = iota
	BLUE
	RED
	YELLOW
	BLACK
	LENTH
)

func get_color (num int) (rcolor Color){
	colors := []Color {WHITE, BLUE, RED, YELLOW, BLACK}

	defer func() {
		if x := recover(); x != nil{
			fmt.Println(x)
			rcolor = BLACK
		}
	}()
	rcolor = colors[num]
	return
}

type Color byte
type age_compare func(p1, p2 person) person
type person struct {
	name string
	age int
	fav_color Color
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
