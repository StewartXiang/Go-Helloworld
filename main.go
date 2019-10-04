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
	//fmt.Println(my_persons)
	var my_persons [][4]human
	c := make(chan ([4]human), 0)
	go prepare_group_async(c)

	get_group_async(c, my_persons)

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
			//fmt.Println(x)
			//fmt.Println("AAAAAAAA")
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
func (p *person) Get_name() (name string){
	name = p.name
	return
}

type human interface {
	Grow()
	Get_age() int
	Get_name() string
}

func Older(p1, p2 human) human{
	println("compare: ", p2.Get_age(), p1.Get_age())
	if p2.Get_age() < p1.Get_age() {
		return p1
	}
	return p2
}

func maker() (my_persons [][2]human){
	var names = [6]string {
		"aaaa", "bbbb", "cccc", "dddd", "eeee", "ffff",
	}
	var ages = [6]int {22, 23, 24, 25, 26, 27}
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
	return my_persons
}

func prepare_group_async(c chan [4]human) {
	var names = [6]string {
		"aaaa", "bbbb", "cccc", "dddd", "eeee", "ffff",
	}
	var names2 = [6]string {
		"aaaa2", "bbbb2", "cccc2", "dddd2", "eeee2", "ffff2",
	}
	var ages = [6]int {22, 23, 24, 25, 26, 27}
	var ages2 = [6]int {222, 232, 242, 252, 262, 272}
	for i, n := range names{
		var ps [4]human
		p := person{
			name:      n,
			age:       ages[i],
			fav_color: get_color(rand.Intn(10)),
		}
		p2 := person{
			name:      names2[i],
			age:       ages2[i],
			fav_color: get_color(rand.Intn(10)),
		}
		ps = [4]human{&p, &p2}
		fmt.Printf("prepare %d\n", i)
		c <- ps
		fmt.Printf("insert over %d\n", i)
		//fmt.Println(ps[0].Get_name())
	}
	close(c)
}

func get_group_async(c chan [4]human, ps [][4]human){
	//这里就可以
	//for {
	//	if p, ok := <-c; ok{
	//		fmt.Printf("get\n")
	//		ps = append(ps, p)
	//	} else {
	//		break
	//	}
	//}
	//这里就不行
	for _ = range c{
		fmt.Printf("get\n")
		p := <- c
		ps = append(ps, p)
	}
	fmt.Println(ps)
	f := filter(ps, Older)
	//fmt.Println(f)
	for _, p := range f{
		fmt.Println("final", p.Get_age(), p.Get_name())
	}
	//for _, p := range f{
	//	fmt.Println(p.Get_age())
	//	p.Grow()
	//	fmt.Println(p.Get_age())
	//}
}

func filter(persons [][4]human, f age_compare) []human{
	var result []human
	for _, group := range persons{
		if value, ok := group[0].(*person); ok {
			fmt.Printf("He is a person: %s\n", value.name)
		} else {
			fmt.Printf("He is a student\n")
		}
		result = append(result, f(group[0], group[1]))
	}
	return result
}

