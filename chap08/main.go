package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//func watchDog(stopch chan bool, task string) {
//	for {
//		select {
//		case <-stopch:
//			fmt.Println("end this goroutin")
//			return
//		default:
//			fmt.Println("I am doing ", task)
//
//		}
//	}
//}
func watchDog(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("i am going to end ", i, "goroutinue")
			return
		default:
			fmt.Println("i am dog", i)
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	ctx, stop := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(num int) {
			defer wg.Done()
			watchDog(ctx, num)
		}(i)
	}
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}
