package main

import (
	"fmt"
	"time"
)

func main()  {
	exit := make(chan int)
	go spiderLauncher(aSpider, 0, exit)
	time.Sleep(10 * time.Second)
	fmt.Printf("I will close launcher!")
	exit <- 1
	time.Sleep(20 * time.Second)
}

func aSpider(n int){
	fmt.Printf("begin spider of %d \n", n)
	time.Sleep(time.Second*3)
	fmt.Printf("finish spider of %d \n", n)
}

func spiderLauncher(spider func(int), n int, exitCh chan int){
	go spider(n)
	timer := time.NewTicker(1 * time.Second)
	select {
	case <- timer.C:
		fmt.Println("after 1s, begin new spider")
		n += 1
		go spiderLauncher(spider, n, exitCh)
	case <- exitCh:
		fmt.Printf("receive exit signal on %d \n", n)
		return
	}
}
