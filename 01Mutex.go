/*
	并发访问场景中不使用锁的例子, 看看出什么问题
	创建10个goroutine,同时不断地对一个变量（count）进行加 1 操作,
	每个 goroutine 负责执行 10 万次的加 1 操作,我们期望的最后计数的结果是 10 * 100000 = 1000000 (一百万).
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	//使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//对变量count执行10万次+1操作
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	//等待10个goroutine完成
	wg.Wait()

	fmt.Println(count)
}

/*
	结果绝对不是100000
	因为，count++ 不是一个原子操作,它至少包含几个步骤,
	比如读取变量 count 的当前值,对这个值加 1,把结果再保存到 count 中.
	因为不是原子操作,就可能有并发的问题

	比如,10 个 goroutine 同时读取到 count 的值为 9527,接着各自按照自己的逻辑加 1,值变成了 9528,
	然后把这个结果再写回到 count 变量.但是,实际上,此时我们增加的总数应该是 10 才对,
	这里却只增加了 1,好多计数都被“吞”掉了.这是并发访问共享数据的常见错误.
*/
