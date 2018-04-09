package models

type HrInterface struct {
	ID 		int64 `json:"id"`
	Cid  	int64 `json:"cid"` // 分类id
	Method 	string `json:"method"` // 接口请求方式
	Name   	string `json:"name"`  // 接口名称
	Url    	string `json:"url"`   // 接口地址
}
