package main

import (
	"fmt"
	"sync"
)

// 一共有 500 张票 现在有2000个人来抢票
func main() {
	// 票数
	ticker := 500 // 表示总共有500张票，这是一个共享资源，多个goroutine会访问并修改它
	// 定义一个互斥锁，用于保护共享资源`ticker`，确保同一时间只有一个goroutine能修改它
	var mu sync.Mutex     // 加一个互斥锁 让顺序变得有序 让一个人抢完了 票数减一 再让下一个人抢
	var wg sync.WaitGroup // 定义一个等待组，用于等待所有goroutine执行完成

	// 2000个人开启2000个goroutine 开抢
	for i := 0; i < 2000; i++ {
		// 这是一个计数器
		wg.Add(1) // 每启动一个goroutine，就增加等待组的计数，表示有一个goroutine需要等待
		// 开启协程
		go func(userId int) { // 使用匿名函数作为goroutine的执行体，传入当前用户的ID
			// 计数器-1
			defer wg.Done() // 在goroutine结束时，调用Done()减少等待组的计数，表示一个goroutine执行完毕
			// 在每一个协程开始之前要 抢占锁
			mu.Lock()         // 获取互斥锁，锁定共享资源，确保同一时间只有一个goroutine能执行到这里
			defer mu.Unlock() // 等到协程执行完以后，再把这个锁释放出去  // 使用defer确保在goroutine结束时释放锁，即使发生panic也会执行
			if ticker > 0 {
				ticker--
				fmt.Printf("用户 %d 购买成功，剩余票数 %d\n", userId, ticker)
			} else {
				fmt.Printf("用户 %d 购买失败\n", userId)
			}
		}(i) // 立即执行匿名函数，并将循环变量i作为参数传入，确保每个goroutine都有唯一的用户ID
	}
	wg.Wait() // 阻塞等待，直到等待组 wg.Add(1) 的计数为0，即所有goroutine都执行完毕
}
