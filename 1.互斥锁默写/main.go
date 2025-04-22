package main

import (
	"fmt"
	"sync"
)

func main() {
	ticker := 500 // 总共有500 张票
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1) // 告诉有几个协程
		go func(userId int) { // 构造协程函数
			// 计数器 -1
			defer wg.Done() // 最后协程都运行完以后 执行这个
			// 加入互斥锁 每一个协程开始之前要抢占锁
			mu.Lock()
			defer mu.Unlock()
			if ticker > 0 {
				ticker--
				fmt.Printf("用户 %d 购买成功，剩余%d\n", userId, ticker)
			} else {
				fmt.Printf("用户 %d 购买失败\n", userId)
			}
		}(i) // 确保每一个协程都有一个用户ID
	}
	wg.Wait() // 阻塞等待，直到wg.Add(1)计数为0 ，表示所有协程执行完毕
}
