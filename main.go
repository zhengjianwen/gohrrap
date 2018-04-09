package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"

	"github.com/zhengjianwen/gohrrap/config"
	"github.com/zhengjianwen/gohrrap/http"
	"github.com/zhengjianwen/gohrrap/models"
)

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	prepare()

	cfg := flag.String("c", "cfg.json", "加载配置文件")
	version := flag.Bool("v", true, "显示版本号")
	help := flag.Bool("h", false, "帮助")
	flag.Parse()

	handleVersion(*version)
	handleHelp(*help)
	handleConfig(*cfg)
}

func main() {
	models.InitMysql()
	go http.Start()

	select {}
}

func handleVersion(displayVersion bool) {
	if displayVersion {
		fmt.Println(config.Version)
		os.Exit(0)
	}
}

func handleHelp(displayHelp bool) {
	if displayHelp {
		flag.Usage()
		os.Exit(0)
	}
}

func handleConfig(configFile string) {
	err := config.Parse(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}
