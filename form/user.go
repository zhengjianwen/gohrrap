package form

import (
	"github.com/zhengjianwen/gohrrap/models"	
	"github.com/zhengjianwen/gohrrap/utils"
)

func UserForm(u *models.HrUser) (bool,string) {
	if u.CnName == "" {
		return false, "用户姓名不能为空"
	}
	if u.Email == "" {
		return false, "邮箱不能为空"
	}
	if !utils.IsEmain(u.Email) || len(u.Email) < 6{
		return false, "邮箱格式不正确"
	}
	if u.Phone == "" {
		return false, "手机号不能为空"
	}
	if u.Username == ""{
		return false,"用户名不能为空"
	}
	if len(u.Username) < 6{
		return false,"用户名长度不能小于6位"
	}
	if !utils.IsAlamNumber(u.Username){
		return false,"用户名只能是英文、数字或下划线"
	}
	if !UserPasswdForm(u.Password){
		return false,"密码验证失败不符合安全要求"
	}
	return true, ""
}

func UserPasswdForm(pasw string) bool {
	if len(pasw) < 6 {
		return false
	}
	return true
}

