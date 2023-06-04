package api

/*

示例代码```
func handler(w http.ResponseWriter, r *http.Request) {
    // 读取请求体
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // 解析请求体中的 JSON 数据
    var reqBody RequestBody
    err = json.Unmarshal(body, &reqBody)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // 处理请求
    respBody := ResponseBody{
        Message: fmt.Sprintf("Hello, %s!", reqBody.Message),
    }

    // 将响应转换为 JSON 格式
    respJSON, err := json.Marshal(respBody)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // 设置响应头
    w.Header().Set("Content-Type", "application/json")

    // 写入响应
    w.Write(respJSON)
}

```

参考 示例代码，实现一个完善的HTTP服务，提供一个GET方法
满足以下要求：
1-创建一个结构体，结构体名为EchoBody，包含以下字段：当前时间、请求方法、请求路径、请求头、请求体等。
2-创建一个GET方法，方法名为'PrintEcho'。
3-为'PrintEcho'方法添加简单注释。
4-返回json格式的状态信息，包括当前时间、请求方法、请求路径、请求头、请求体等。
*/

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type EchoBody struct {
	CurrentTime time.Time `json:"current_time"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	Header      string    `json:"header"`
	Body        string    `json:"body"`
}

// @Summary PrintEcho
// @Description Print Request Info in JSON, include current time, request method, request path, request header, request body
// @Success 200 {string} string	"ok"
// @Router /api/v1/echo [GET][个人测试]
func PrintEcho(w http.ResponseWriter, r *http.Request) {
	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 处理请求
	respBody := EchoBody{
		CurrentTime: time.Now(),
		Method:      r.Method,
		Path:        r.URL.Path,
		Header:      r.Header.Get("Content-Type"),
		Body:        string(body),
	}

	// 将响应转换为 JSON 格式
	respJSON, err := json.Marshal(respBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 写入响应
	w.Write(respJSON)
}
