package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	n := make(chan int)

	go PrintGo(c, n)

	tiomeout := time.After(4 * time.Second)
	fmt.Println("ждём сообщение в канал")
	select {
	case r := <-c:
		fmt.Printf("получено чётное значение r = %d \n", r)
	case r := <-n:
		fmt.Printf("получено нечётное значение r = %d \n", r)
	case <-tiomeout:
		fmt.Println("функция выполняется слишком долго, я не могу столько ждать")
		return
	}

	fmt.Println("main отработала")
}

func PrintGo(chet, nechet chan int) {
	time.Sleep(3 * time.Second)

	i := rand.Intn(100)
	if i%2 == 0 {
		chet <- i
	} else {
		nechet <- i
	}

	fmt.Println("горутина 1 отработала")
}
