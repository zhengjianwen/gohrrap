package http

import (
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/toolkits/file"

	"github.com/zhengjianwen/gohrrap/config"
	"github.com/zhengjianwen/gohrrap/handler"

	"github.com/zhengjianwen/gohrrap/http/middleware"
	"github.com/zhengjianwen/gohrrap/render"
)

func Start() {
	render.Init()
	//context.Init()
	// router
	r := mux.NewRouter().StrictSlash(false)
	handler.ConfigRouter(r)
	// 中间件
	n := negroni.New()
	n.Use(middleware.NewRecovery(file.MustOpenLogFile(config.G.Log.Error)))
	n.Use(middleware.NewLogger(file.MustOpenLogFile(config.G.Log.Access)))

	n.UseHandler(r)

	log.Println("listening on", config.G.Http.Listen)
	log.Fatal(http.ListenAndServe(config.G.Http.Listen, n))
}
