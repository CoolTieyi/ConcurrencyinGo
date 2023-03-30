/*
	Mutex 的基本用法
	进入临界区前加锁, 离开临界区后释放锁

*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	//互斥锁保护计数器
	var mu sync.Mutex
	//计数器
	var count = 0

	//辅助变量, 确认所有goroutine都能完成
	//A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(10)

	//启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			//Done decrements the WaitGroup counter by one.
			defer wg.Done()
			//累加10w次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	//Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
	fmt.Println(count)
}
