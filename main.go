package main

/*

示例代码```

func main() {
    // 打印日志
    log.Println("服务已启动")

    // 注册 HTTP 处理函数
    http.HandleFunc("/api/v1/Service", Service)

    // 创建一个信号通道
    sigChan := make(chan os.Signal, 1)

    // 监听 SIGINT 信号
    signal.Notify(sigChan, syscall.SIGINT)

    // 启动 HTTP 服务
    go func() {
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
            log.Fatal("启动服务失败：", err)
        }
    }()

    // 等待 SIGINT 信号
    <-sigChan

    // 打印日志
    log.Println("用户手动停止服务")
}

```

参考 示例代码，完成main()方法，实现一个简单的HTTP服务，监听8080端口，完成以下目标：
1-引入"testapi/api"的HealthCheck方法，提供/api/v1/healthcheck接口。
2-引入“testapi/api”的PrintEcho方法，提供/api/v1/echo接口。
3-当服务启动时，打印日志：Server is running on port 8080
4-当服务收到请求时，打印日志：Request received
5-当服务启动失败时，打印日志：Failed to start server 以及错误信息
6-当服务停止时，打印日志：Server stopped
7-当用户手动停止服务时，打印日志：User terminated
*/

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testapi/api"
)

func main() {
	// 打印日志
	log.Println("服务已启动")

	// 注册 HTTP 处理函数
	http.HandleFunc("/api/v1/healthcheck", api.HealthCheck)
	http.HandleFunc("/api/v1/echo", api.PrintEcho)
	http.HandleFunc("/api/v1/index", api.GetIndexName)
	http.HandleFunc("/api/v1/query", api.QueryIndex)

	// 创建一个信号通道
	sigChan := make(chan os.Signal, 1)

	// 监听 SIGINT 信号
	signal.Notify(sigChan, syscall.SIGINT)

	// 启动 HTTP 服务
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("启动服务失败：", err)
		}
	}()

	// 等待 SIGINT 信号
	<-sigChan

	// 打印日志
	log.Println("用户手动停止服务")
}
