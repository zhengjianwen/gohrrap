package handler

import (
	"fmt"
	"net/http"
	"github.com/zhengjianwen/gohrrap/config"
)

func Version (w http.ResponseWriter, r *http.Request)  {

	fmt.Println("[Version]查看版本:",config.Version)
	resp :=ResponseData{Status:true,Message:config.Version,Id:config.Version}
	RespDate(w,resp)
}
