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
	var names = [6]string {
		"aaaa", "bbbb", "cccc", "dddd", "eeee", "ffff",
	}
	var ages = [6]int {22, 23, 24, 25, 26, 27}
	//p := person{
	//	name:      "test",
	//	age:       10,
	//	fav_color: RED,
	//}
	//p.Grow()
	//s := student{
	//	person:     p,
	//	student_id: "none",
	//}
	//s.Grow()
	var my_persons [][2] human
	for i, n := range names{

		p := person{
			name:      n,
			age:       ages[i],
			fav_color: get_color(rand.Intn(10)),
		}
		if i%2 == 0{
			my_persons = append(
				my_persons,
				[2]human{&p})
		}else{
			my_persons[len(my_persons)-1][1] = &p
		}
	}
	my_persons = append(
		my_persons,
		[2]human{
			&student{person{
				name:      "gggg",
				age:       28,
				fav_color: RED,
			}, "std1"},
			&student{person{
				name:      "hhhh",
				age:       29,
				fav_color: BLUE,
			}, "std2"},
		})
	//fmt.Println(my_persons)
	f := filter(my_persons, Older)
	for _, p := range f{
		fmt.Println(p.Get_age())
	}
}


const(
	WHITE = iota
	BLUE
	RED
	YELLOW
	BLACK
)

func get_color (num int) (rcolor Color){
	colors := []Color {WHITE, BLUE, RED, YELLOW, BLACK}

	defer func() {
		if x := recover(); x != nil{
			fmt.Println(x)
			fmt.Println("AAAAAAAA")
			rcolor = BLACK
		}
	}()
	rcolor = colors[num]
	return
}

type Color byte
type age_compare func(p1, p2 human) human
type person struct {
	name string
	age int
	fav_color Color
}
type student struct {
	person
	student_id string
}

func (p *person) Grow(){
	p.age += 1
}
func (p *person) Get_age() (age int){
	age = p.age
	return
}

type human interface {
	Grow()
	Get_age() int
}

func Older(p1, p2 human) human{
	if p1.Get_age()>p2.Get_age() {
		return p1
	}
	return p2
}

func filter(persons [][2]human, f age_compare) []human{
	var result []human
	for _, group := range persons{
		result = append(result, f(group[0], group[1]))
	}
	return result
}
