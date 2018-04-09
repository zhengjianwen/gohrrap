package models

import (
	"github.com/zhengjianwen/gohrrap/utils"
	"github.com/zhengjianwen/gohrrap/form"
	"github.com/zhengjianwen/gohrrap/log"
)

type HrUser struct {
	ID 			int64 `json:"id"`
	Username   	string `json:"username" xorm:"notnull 'username'"`  // 用户名称
	Password   	string `json:"-"`
	Phone 	   	string `json:"phone" xorm:"notnull 'phone'"`
	Email 		string `json:"email" xorm:"notnull 'email'"`
	Qq          string `json:"qq" xorm:"null 'qq'"`
	CnName     	string `json:"cn_name" xorm:"notnull 'cnname'"`
	Status      int64 `json:"status" `   // 1 激活  0 冻结
}

var HrUserRepo  = new(HrUser)

func (u *HrUser)Save() (int64,string) {
	u.Status = 1
	u.Password = utils.EncryptPassword(u.Password)
	status,msg := form.UserForm(u)
	if status{
		ishasUsername(u)
		_,err := Orm.Insert(u)
		if err != nil{
			return 0,"注册失败"
		}
		return u.ID,""
	}
	return 0,msg

}

func ishasUsername(u *HrUser) bool {
	bean := new(HrUser)
	has,err := Orm.Where("username = ? or phone = ? or email =?",u.Username).Get(bean)
	if err != nil{
		log.Errorf("查询用户名是否重复错误：%v",err)
		return false
	}
	if has{
		return true
	}

	return false
}