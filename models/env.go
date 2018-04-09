package models

type HrEnv struct {
	Id    int64 `json:"id"`
	Pid   int64 `json:"pid"` // 项目ID
	Name  string `json:"name"`
	Url   string `json:"url"`
}
