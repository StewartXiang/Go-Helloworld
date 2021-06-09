package main

import (
	"fmt"
	"time"
)

func main()  {
	go spiderLauncher(aSpider, 0)
	time.Sleep(60 * time.Second)

}
func aSpider(n int){
	fmt.Printf("begin spider of %d \n", n)
	time.Sleep(time.Second*3)
	fmt.Printf("finish spider of %d \n", n)
}

func spiderLauncher(spider func(int), n int){
	go spider(n)
	timer := time.NewTicker(1 * time.Second)
	select {
	case <- timer.C:
		fmt.Println("after 1s, begin new spider")
		n += 1
		go spiderLauncher(spider, n)
	}
}
