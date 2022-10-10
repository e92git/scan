package main

import (
	"fmt"
	"sync"
)

// go run playground/mutexGroup.go
// go run -race ./playground/mutexGroup.go
func main() {

    // greate work
	c := 0
	n := 200
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			m.Lock()
			c++
			m.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(c)

    // channel
	cc := 0
	nn := 2
	chanWg := sync.WaitGroup{}
	chanWg.Add(1)
	ch := make(chan struct{}, nn)
	go func() {
		for i := 1; ; i++ {
			select { // блокируемся, пока в один из каналов не попадёт сообщение
			case <-ch:
                cc++
				println("channel 1", i, cc)
			}
			if i >= 2 { // через 2 итерации выходим (иначе будет deadlock)
				break
			}
		}
		fmt.Println("ch", ch)
		// // for range ch {
		// 	cc++
		// 	fmt.Println("CCC", cc)
		// // }
		// fmt.Println("sleep", cc)
		chanWg.Done()
	}()
	ch <- struct{}{}
	ch <- struct{}{}
	fmt.Println("Step 1", cc)
	chanWg.Wait()
	fmt.Println("Step 2", cc)
	fmt.Println("Done")
}
