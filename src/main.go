package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置路由和处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
	})

	// 启动服务器并监听指定端口
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}