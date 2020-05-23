package dbtoapi

import "encoding/json"

//http响应
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//包装为http响应，返回字符串
func respJson(data interface{}) string {
	resp := Response{
		Status:  200,
		Message: "success",
		Data:    data,
	}
	respJson, _ := json.Marshal(resp)
	respJsonStr := string(respJson)
	return respJsonStr
}
