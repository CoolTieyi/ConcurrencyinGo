/*
	Nutex的四个阶段: 初版--给新人机会--多给些机会--解决饥饿

	初版Mutex:
		使用flag表示锁是否被持有
	给新人机会:
		让新的goroutine也尽可能地先获得锁
	多给些机会:
		照顾新来的和被唤醒的goroutine, 但会饥饿
	解决饥饿:
		解决饥饿问题
*/

//2008年的第一版Mutex
package main

/*
CAS:Compare And Swap 比较并交换
	CAS指令将一个内存地址中的值和预期的值进行比较, 如果是地址的值等于预期的值,
	就将内存地址中的值替换为新值, 这个操作是原子性的.
	【说白了就是:V和A相等的话，就把V换成B,是原子性的，所以一次就一个线程成功】
	CAS是实现互斥锁和同步原语的基础
原子性：
	原子性保证这个指令总是基于最新的值进行计算, 如果同时有其他县城修改了这个值,
	那么CAS会返回失败
*/
func cas(val *int32, old, new int32) bool
func semacquire(*int32)
func semrelease(*int32)

//互斥锁的结构, 包含两个字段
type Mutex struct {
	key  int32 //是一个flag,标识锁是否被goroutine持有; 0未持有\1被持有无等待者\n被持有同时有n-1个等待者
	sema int32 //信号量专用,用以控制goroutine阻塞或唤醒
}

//保证成功在val上增加delta的值
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) { //这里期望的值是val = v+delta
			return v + delta
		}
	}
	panic("unreached")
}

//请求锁
func (m *Mutex) Lock() {
	if xadd(&m.key, 1) == 1 { //标识+1，如果等于1，就成功获取到锁
		return
	}
	semacquire(&m.sema) //否则阻塞等待, 使用信号量将自己休眠,等锁释放的时候, 信号量会将它唤醒
}

func (m *Mutex) Unlock() {
	if xadd(&m.key, -1) == 0 { //标识-1，如果等于0，则没有其他进程等待
		return
	}
	semrelease(&m.sema) //唤醒其他阻塞的goroutine
}

/*
!!!!!!!!!!!问题在于:
1. 	Unlock()可以被任意的goroutine调用释放锁, 即便是没有持有这个互斥锁的gotoutine, 也可以进行操作
	因为Mutex本身并没有包含持有这把锁的goroutine的信息, 所以Unlock不会对此检查,Mutex这个设计保持至今
2.  请求锁的 goroutine 会排队等待获取互斥锁。虽然这貌似很公平 但是从性能上来看 却不是最优的
	因为如果我们能够把锁交给正在占用 CPU 时间片的 goroutine 的话 那就不需要做上下文的切换
	在高并发的情况下 可能会有更好的性能  =====> 第二版:给新人机会
*/
