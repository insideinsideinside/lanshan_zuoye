package main

import (
	"fmt"
)

var (
	Wallet = 0 // 一贫如洗的泡泡
)

func main() {
	get := make(chan struct{}, 1)

	for i := 0; i <= 10_000; i++ { // 泡泡成功骗取到了 1w 个人的同情
		get <- struct{}{}
		go func() {
			<-get
			//go vPaopao50()
			Wallet += 50

		}()

	}
	//time.Sleep(2 * time.Second) // 可恶的泡泡竟然睡起了大觉
	fmt.Println("泡泡现在有", Wallet, "元")
	// 睡醒的泡泡真的可以获得他乞讨到的 1w * 50 = 50w 元么？
}

//func vPaopao50() {
//
//	Wallet += 50
//
//}
