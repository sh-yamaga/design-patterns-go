package main

import (
	"fmt"
	"sync"
)

// Singleton構造体
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// GetInstanceはSingletonのインスタンスを返します
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "初期データ"}
	})
	return instance
}

func main() {
	// シングルトンインスタンスを取得
	s1 := GetInstance()
	fmt.Println(s1.data)

	// 別の場所で再度シングルトンインスタンスを取得
	s2 := GetInstance()
	fmt.Println(s2.data)

	// 同じインスタンスであることを確認
	if s1 == s2 {
		fmt.Println("s1とs2は同じインスタンスです。")
	}
}
