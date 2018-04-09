package models

import "time"

// 用户登录历史
type HrLoginHistory struct {
	ID  		int64 `json:"id" xorm:"pk autoincr 'id'"`
	Uid  		int64 `json:"uid" xorm:"'uid'"`
	LoginIp 	string `json:"login_ip" xorm:"'login_ip'"` // 登录IP
	LoginAddr   string `json:"login_addr" xorm:"'login_addr'"` // 登录地址
	LoginTime   time.Time `json:"login_time" xorm:"created 'login_time'"` // 登录时间
}
