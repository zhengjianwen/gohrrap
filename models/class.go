package models

// 接口分类
type HrClass struct {
	Id    int64 `json:"id"`
	Pid   int64 `json:"pid"`
	Name  string `json:"name"`
} 