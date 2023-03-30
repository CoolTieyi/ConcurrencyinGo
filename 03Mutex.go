/*
	将Mutex写进结构体	, 可以匿名调用
*/
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {
	var counter Counter
	//A WaitGroup waits for a collection of goroutines to finish
	var wg sync.WaitGroup
	//Add goroutines
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			//关闭
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
