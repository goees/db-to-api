package dbtoapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func RunServer(config *Config) {
	conf = config
	http.HandleFunc("/tables", apiTables())
	http.HandleFunc("/query", apiQuery())
	log.Fatal(http.ListenAndServe(":"+conf.HttpServerPort, nil))
}

//api, "/tables": 查询有哪些表
func apiTables() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//返回
		resp := Response{
			Status:  200,
			Message: "Success",
			Data:    nil,
		}
		//查询有哪些表
		result := QueryDBTables()
		resp.Data = result
		//响应
		response(w, resp)
	}
}

//api，"/query"：查询指定表
func apiQuery() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//返回
		resp := Response{
			Status:  200,
			Message: "Success",
			Data:    nil,
		}
		//校验
		table, pageNum, pageSize := checkRequest(r, &resp)
		//校验通过，则开始查询
		if resp.Status == 200 {
			result := QueryDB(table, pageNum, pageSize)
			resp.Data = result
		}
		//响应
		response(w, resp)
	}
}

//http响应
func response(w http.ResponseWriter, resp Response) {
	//响应
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	respJson, _ := json.Marshal(resp)
	respJsonStr := string(respJson)
	_, errR := fmt.Fprintf(w, respJsonStr)
	checkErr(errR)
}

//api "/query" 的参数校验
func checkRequest(r *http.Request, resp *Response) (table string, pageNum, pageSize int) {
	if http.MethodPost != r.Method {
		resp.Status = 206
		resp.Message = "请使用POST请求"
	}

	//获取参数：table、page、size
	body, _ := ioutil.ReadAll(r.Body)
	//存储参数的map
	requestBody := make(map[string]string)
	//json解析为map
	err := json.Unmarshal(body, &requestBody)
	if err != nil {
		resp.Status = 206
		resp.Message = "json解析异常"
	}

	//校验map
	table = requestBody["table"]
	page := requestBody["page"]
	size := requestBody["size"]
	if len(table) == 0 {
		//没有table参数
		resp.Status = 206
		resp.Message = "缺少table参数"
	}
	if len(page) == 0 {
		page = "1"
	}
	if len(size) == 0 {
		size = "10"
	}
	pageNum, errA := strconv.Atoi(page)
	pageSize, errB := strconv.Atoi(size)
	if errA != nil || errB != nil {
		//page和size参数必须是数字
		resp.Status = 206
		resp.Message = "page和size参数必须是数字"
	}
	return table, pageNum, pageSize
}
