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
1-创建一个结构体，结构体名为'HealthBody'，包含以下字段：当前时间、服务名称、服务版本、服务状态。
2-创建一个GET方法，方法名为'HealthCheck'。
3-为'HealthCheck()'方法添加简单注释。
3-返回json格式的状态信息，包括当前时间、服务名称、服务版本、服务状态等。
*/

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthBody is a struct
type HealthBody struct {
	CurrentTime time.Time `json:"currentTime"`
	ServiceName string    `json:"serviceName"`
	Version     string    `json:"version"`
	Status      string    `json:"status"`
}

// HealthCheck
// @Summary HealthCheck
// @Description HealthCheck
// @Tags HealthCheck
// @Accept json
// @Produce json
// @Success 200 {object} HealthBody
// @Router /api/v1/healthcheck [GET][个人测试]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthBody := HealthBody{
		CurrentTime: time.Now(),
		ServiceName: "testapi",
		Version:     "v1",
		Status:      "ok",
	}
	healthBodyJSON, err := json.Marshal(healthBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(healthBodyJSON)
}
