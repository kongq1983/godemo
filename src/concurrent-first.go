package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()的函数的并发goroutine
	go producer("数据1", channel)
	go producer("数据2", channel)

	// 数据消费函数
	customer(channel)
}

// 生产者
func producer(header string, channel chan<- string) {

	// 无限循环，不停的生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1s
		time.Sleep(time.Second)
	}

}

func customer(channel <-chan string) {
	// 不停的获取数据
	for {
		//从通道中取出数据，此处会阻塞通道中返回数据
		message := <-channel
		// 打印数据
		fmt.Printf("message=%s \n", message)
	}

}
