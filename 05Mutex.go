/*
	第二版Mutex --2011.6.30
*/
package main

type Mutex struct {
	state int32
	sema  uint32
} /*state字段被分为3个部分, 代表三个数据
第一位 mutexWaiters 阻塞等待的waiter数量
第二位 mutexWoken 唤醒标记
第三位 mutexLocked 持有锁的标记
*/

const (
	mutexLocketed = 1 << iota //	mutex is locked
	mutexWoken
	mutexWaiterShift = iota
)

//请求锁的方法
func (m *Mutex) Lock() {
	//Fast Path: 幸运case, 能直接获取到锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
	awoke := false
	for {
		old := m.state
		new := old | mutexLocked //新状态加锁
		if old&mutexLocked != 0 {
			new = old + 1<<mutexWaiterShift //等待者数量加以
		}
		if awoke {
			//goroutine是被唤醒着的
			//新状态清除唤醒标志
			new &^= mutexWoken
		}
		//设置新状态
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&mutexLocked == 0 { //锁原状态未加锁
				break
			}
			runtime.Semacquire(&m.sema) //请求信号量
			awoke = true
		}
	}
}
