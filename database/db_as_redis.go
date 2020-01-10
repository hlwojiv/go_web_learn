package main

import (
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"

	// 字符串操作
	// client.Set("b", []byte("hello"))
	// val, _ := client.Get("a")
	// fmt.Println(string(val))

	// list 操作
	vals := []string{"a", "b", "c", "d", "e"}
	for _, v := range vals {
		client.Rpush("l", []byte(v))
	}
	dbvals, _ := client.Lrange("l", 0, 4)
	for i, v := range dbvals {
		println(i, "：", string(v))
	}
}
