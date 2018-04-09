package models

import (

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/zhengjianwen/gohrrap/config"
	"fmt"
)

var Orm *xorm.Engine

func InitMysql() {
	if config.G.Mysql.Addr != ""{
		addr := fmt.Sprintf("%s/hrrap?charset=utf8&loc=Asia%%2FShanghai",config.G.Mysql.Addr)
		db, err := xorm.NewEngine("mysql", addr)
		if err != nil {
			log.Fatalf("[InitMysql] connect mysql => %s\n", addr)
		}
		db.SetMaxIdleConns(config.G.Mysql.Idle)
		db.SetMaxOpenConns(config.G.Mysql.Max)
		db.ShowSQL(config.G.Mysql.ShowSQL)
		db.Logger().SetLevel(core.LOG_INFO)
		// 链接Cndb 原始数据表
		db.Exec("create database IF NOT EXISTS 'sys' character set utf8;")

	}else {
		log.Print("[sys] 无法连接数据库")
	}


}

func ConnOrm(orgid int64) (*xorm.Engine,error){
	addr := fmt.Sprintf("%s/cmdb_org_%d?charset=utf8&loc=Asia%%2FShanghai",config.G.Mysql.Addr,orgid)
	if config.G.Mysql.Addr != ""{
		db, err := xorm.NewEngine("mysql", addr)
		if err != nil {
			log.Fatalf("[InitMysql] connect mysql => %s\n", config.G.Mysql.Addr)
		}
		db.ShowSQL(config.G.Mysql.ShowSQL)
		db.Logger().SetLevel(core.LOG_INFO)
		return db,nil
	}else {
		log.Print("[sys] 无法连接数据库",addr)
		return nil,fmt.Errorf("无法连接数据库")
	}
}