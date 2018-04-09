package middleware

import (
	"net/http"

	"github.com/zhengjianwen/gohrrap/http/cookies"
	"github.com/zhengjianwen/gohrrap/render"
	"github.com/toolkits/web/errors"
)

func Authentication(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userid, username := cookie.ReadUser(r)
	if userid <= 0 {
		panic(errors.NotLoginError())
	}

	render.Put(r, "userid", userid)
	render.Put(r, "username", username)

	next(rw, r)
}
