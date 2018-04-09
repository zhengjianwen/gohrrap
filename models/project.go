package models

import "time"

type HrProject struct {
	Id   int64 `json:"id"`
	Name  string `json:"name"`   // 项目名称
	Desc  string `json:"desc"`  // 描述
	Ctime  time.Time `json:"ctime" xorm:"created 'ctime'"` // 更新时间
	Uptime  time.Time `json:"uptime" xorm:"created 'uptime'"` // 创建时间
}
