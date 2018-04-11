package models

// 权限  1 查  2 增 4改  8删
// 项目权限  修改  删除  授权
// 分类权限
// 接口权限  创建  修改  删除
// 用户权限  添加2  移除8
// 产品设计  创建  修改  删除
// 字典功能  增加  修改  删除
//

type HrMember struct {
	ID   		int64 	`json:"id"`
	Uid  		int64 	`json:"uid"`   // 用户ID
	Pid  		int64 	`json:"pid"`   // 项目ID
	Permission 	string 	`json:"permission"` // 权限 pro项目 interface 接口  user成员
	Operation   int64 	`json:"operation"`


}