package main

import (
	"context"
	"time"
)

func main() {
	ctx1, cancel := context.WithCancel(context.TODO())

	// 开启新协程模拟新的任务，并且引用主线程的ctx
	go func() {
		// 新的协程会监听ctx是否超时，一旦超时就return
		println("运行的业务逻辑")
		//for true {
		//	time.Sleep(500*time.Millisecond)
		//	println("1111111")
		//}

		select {
		case <-ctx1.Done():
			println("子线程：我引用的ctx已超时！我只能中断任务，直接返回")
			return
		}
	}()

	time.Sleep(time.Second)

	println("主线程调用cancel模拟主线程结束，此时子线程会通过ctx.Done得到通知")
	cancel()

	time.Sleep(time.Second) // 稍等等待一下，可以看到子线程打印信息
}