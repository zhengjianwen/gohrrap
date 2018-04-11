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
	IsAdmin     int64 `json:"is_admin" xorm:"'is_admin'"`
	Status      int64 `json:"status" `   // 1 激活  0 冻结
}

var HrUserRepo  = new(HrUser)

func (u *HrUser)Save() (int64,string) {
	u.Status = 1
	status,msg := form.UserForm(u)
	u.Password = utils.EncryptPassword(u.Password)
	if status{
		if ishasUsername(u){
			return 0,"手机号、邮箱或用户名已存在,请修改后重试"
		}
		_,err := Orm.Insert(u)
		if err != nil{
			log.Errorf("用户注册失败，%v",err)
			return 0,"注册失败"
		}
		return u.ID,""
	}
	return 0,msg

}

func (u *HrUser)Update() (string) {
	_,err := Orm.Where(u.ID).Cols("qq","email","phone").Update(u)
	if err != nil{
		log.Errorf("用户更新失败，%v",err)
		return "删除失败"
	}
	return ""
}

func (u *HrUser)FreezeUser() (string) {
	u.Status = 0
	_,err := Orm.Where(u.ID).Cols("status").Update(u)
	if err != nil{
		log.Errorf("用户冻结失败，%v",err)
		return "冻结失败"
	}
	return ""
}

// 修改密码
func (u *HrUser)UpdatePassworf() (string) {
	if !form.UserPasswdForm(u.Password){
		return "密码不符合要求"
	}
	u.Password = utils.EncryptPassword(u.Password)
	_,err := Orm.Where(u.ID).Cols("status").Update(u)
	if err != nil{
		log.Errorf("用户冻结失败，%v",err)
		return "冻结失败"
	}
	return ""
}

func (u *HrUser)UndoUser() (string) {
	u.Status = 1
	_,err := Orm.Where(u.ID).Cols("status").Update(u)
	if err != nil{
		log.Errorf("用户解封失败，%v",err)
		return "解封失败"
	}
	return ""
}

func (u *HrUser)Del() (string) {
	_,err := Orm.Where(u.ID).Delete(u)
	if err != nil{
		log.Errorf("用户注册失败，%v",err)
		return "删除失败"
	}
	return ""
}



func ishasUsername(u *HrUser) bool {
	bean := new(HrUser)
	has,err := Orm.Where("username = ? or phone = ? or email =?",u.Username,u.Phone,u.Email).Count(bean)
	if err != nil{
		log.Errorf("查询用户名是否重复错误：%v",err)
		return false
	}
	if has > 0{
		return true
	}

	return false
}