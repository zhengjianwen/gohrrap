package models


// 请求头部信息
type HrHeader struct {
	ID 		int64 `json:"id"`
	Iid  	int64 `json:"iid"` // 接口id
	Key 	string `json:"key"` // 键
	Vaule   string `json:"vaule"`  // 值
	Desc    string `json:"desc"`   // 描述
}
