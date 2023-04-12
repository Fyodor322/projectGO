package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go PrintGo(c)
	fmt.Println("main отработала")

	fmt.Println("ждём сообщение в канал")
	r := <-c
	fmt.Println(r)
}

func PrintGo(c chan int) {
	time.Sleep(3 * time.Second)
	c <- 100
	fmt.Println("горутина отработала")
}
