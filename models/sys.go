package models

import (
	"github.com/zhengjianwen/gohrrap/log"
)
// 系统参数
type HrSys struct {
	Id    			int64 `json:"id"`
	WebName  		string 	`json:"web_name" xorm:"'web_name'"`		// 网站名称
	Keywords  		string 	`json:"keywords" xorm:"'keywords'"`				// 网站关键字
	Description  	string 	`json:"description" xorm:"'description'"`   // 网站描述
	Copyright  		string 	`json:"copyright" xorm:"'copyright'"`   // 版权信息
	Email  			string 	`json:"email" xorm:"'email'"`   // 版权信息
	RegisterToken  	string `json:"register_token" xorm:"'register_token'"`   // 注册口令,邀请码
	RegisterCaptcha int8 	`json:"register_captcha" xorm:"'register_captcha'"`   // 注册验证码开关
	LoginCaptcha 	int8 	`json:"login_captcha" xorm:"'login_captcha'"`   // 登录验证码开关
}

var HrSysrepo = new(HrSys)

func InitSys() error {
	bean := new(HrSys)
	bean.Id = 1
	bean.WebName = "海瑞Api接口管理平台"
	bean.Keywords = "海瑞Api接口管理平台|海瑞"
	bean.Description = "海瑞Api接口管理平台是一个API接口文档管理系统，致力于减少前后端沟通成本，提高团队协作开发效率"
	bean.Copyright = "Copyright ©2018 hairuinet.com"
	bean.Email = "574601624@qq.com"
	bean.RegisterToken = "hrrap"
	bean.RegisterCaptcha = 1
	bean.LoginCaptcha = 1
	_,err := Orm.Insert(bean)
	if err != nil{
		log.Errorf("初始化公司失败，错误：%v",err)
	}
	return err
}

func (sys *HrSys)Update() error {
	_,err := Orm.Id(1).Update(sys)
	if err != nil{
		log.Errorf("更新系统信息错误：%v",err)
	}
	return err
}