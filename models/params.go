package models


// 请求参数 form 请求 键值对方式
type HrParams struct {
	ID 		int64 `json:"id"`
	Iid  	int64 `json:"iid"` // 接口id
	Name 	string `json:"name"` // 字段别名
	Title   string `json:"title"`  // 字段含义
	Type    int64 `json:"type"`   // 字段类型 1 string 2 int 3 float 4 bool 5 arr
	Must    int64 `json:"must"`  // 是否必填
	Default string `json:"default"` // 默认值
	Desc    string `json:"desc"` // 备注说明
}

