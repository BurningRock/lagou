package main

import (
	"fmt"
	"sync"
	"time"
)

func watchDog(stopch chan bool, task string) {
	for {
		select {
		case <-stopch:
			fmt.Println("end this goroutin")
			return
		default:
			fmt.Println("I am doing ", task)

		}
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan bool)
	go func() {
		defer wg.Done()
		go watchDog(ch, "监控狗1")
	}()
	time.Sleep(5 * time.Second)
	ch <- true
	fmt.Println("after 5 seconds ,the task is over")
	wg.Wait()
}
