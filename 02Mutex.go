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
		/*
			go func(){}() 是 Go 语言中的一种语法, 它表示创建一个匿名函数并立即在一个新的 goroutine 中执行它.
			go 关键字表示在一个新的 goroutine 中执行函数, func(){} 表示定义一个匿名函数,最后的 () 表示立即调用这个匿名函数.
		*/
		go func() {
			//Done decrements the WaitGroup counter by one. Done函数将WaitGroup计数器减一
			defer wg.Done()
			//累加10w次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	//Wait blocks until the WaitGroup counter is zero. Wait 函数会阻塞,直到 WaitGroup 计数器变为零.
	wg.Wait()
	fmt.Println(count)
}
