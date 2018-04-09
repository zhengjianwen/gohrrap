package models

type HrRaw struct {
	ID 		int64 `json:"id"`
	Iid  	int64 `json:"iid"` // 接口id
	Name 	string `json:"name"` // 字段别名
	Title   string `json:"title"`  // 字段含义
	Type    int64 `json:"type"`   // 字段类型 1 arr 2 dic
	Must    int64 `json:"must"`  // 是否必填
	Default string `json:"default"` // 默认值
	Desc    string `json:"desc"` // 备注说明
}

