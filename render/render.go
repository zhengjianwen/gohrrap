package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/unrolled/render"

	"github.com/zhengjianwen/gohrrap/config"
)

var Render *render.Render
var fm = template.FuncMap{
	"safe": func(raw string) template.HTML {
		return template.HTML(raw)
	},
}

func Init() {
	debug := config.G.Debug
	Render = render.New(render.Options{
		Directory:     "views",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"{{", "}}"},
		IndentJSON:    false,
		Funcs:         []template.FuncMap{fm},
		IsDevelopment: debug,
	})
}

func Put(r *http.Request, key string, val interface{}) {
	m, ok := context.GetOk(r, "DATA_MAP")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		context.Set(r, "DATA_MAP", mm)
	} else {
		context.Set(r, "DATA_MAP", map[string]interface{}{key: val})
	}
}

func HTML(r *http.Request, w http.ResponseWriter, name string, htmlOpt ...render.HTMLOptions) {
	Render.HTML(w, http.StatusOK, name, context.Get(r, "DATA_MAP"), htmlOpt...)
}

func Error(w http.ResponseWriter, err error) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}

	Render.JSON(w, http.StatusOK, map[string]string{"msg": msg})
}

func Message(w http.ResponseWriter, format string, args ...interface{}) {
	Render.JSON(w, http.StatusOK, map[string]string{"msg": fmt.Sprintf(format, args...)})
}

func Data(w http.ResponseWriter, v interface{}, err ...error) {
	Render.JSON(w, http.StatusOK, v)
}

func Text(w http.ResponseWriter, v string) {
	Render.Text(w, http.StatusOK, v)
}

func Respone(w http.ResponseWriter, v interface{}) {
	Render.JSON(w, http.StatusOK, v)
}

func ResponeFile(w http.ResponseWriter, file []byte, filename string,v interface{}) {
	if len(file) == 0{
		Respone(w,v)
	}else {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
		Render.Data(w, 200, file)
	}
}
