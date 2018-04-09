package cookie

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/zhengjianwen/gohrrap/config"
	"fmt"
)

const USER_COOKIE_NAME = "hrrap"

var SecureCookie *securecookie.SecureCookie

func Init() {
	var hashKey = []byte(config.G.Http.Secret)
	var blockKey = []byte(nil)
	SecureCookie = securecookie.New(hashKey, blockKey)
}

type CookieData struct {
	UserId   int64
	Username string
}

func ReadUser(r *http.Request) (int64, string) {
	if cookie, err := r.Cookie(USER_COOKIE_NAME); err == nil {
		var value CookieData
		if err = SecureCookie.Decode(USER_COOKIE_NAME, cookie.Value, &value); err == nil {
			return value.UserId, value.Username
		}
	}

	return 0, ""
}

func WriteUser(w http.ResponseWriter, id int64, username string) error {
	value := CookieData{UserId: id, Username: username}
	encoded, err := SecureCookie.Encode(USER_COOKIE_NAME, value)
	if err != nil {
		return fmt.Errorf("写入Cookies 失败")
	}

	var cookie *http.Cookie
	if config.G.Domain != ""{
		cookie = &http.Cookie{
			Name:     USER_COOKIE_NAME,
			Value:    encoded,
			Path:     "/",
			MaxAge:   3600 * 2,
			HttpOnly: true,
			Domain:config.G.Domain,
		}
	}else {
		cookie = &http.Cookie{
			Name:     USER_COOKIE_NAME,
			Value:    encoded,
			Path:     "/",
			MaxAge:   3600 * 2, // 2小时
			HttpOnly: true,
		}
	}

	http.SetCookie(w, cookie)
	return nil
}

func RemoveUser(w http.ResponseWriter) error {
	var value CookieData
	encoded, err := SecureCookie.Encode(USER_COOKIE_NAME, value)
	if err != nil {
		return err
	}

	var cookie *http.Cookie
	if config.G.Domain != ""{
		cookie = &http.Cookie{
			Name:     USER_COOKIE_NAME,
			Value:    encoded,
			Path:     "/",
			MaxAge:   60,
			HttpOnly: true,
			Domain:config.G.Domain,
		}
	}else {
		cookie = &http.Cookie{
			Name:     USER_COOKIE_NAME,
			Value:    encoded,
			Path:     "/",
			MaxAge:   60,
			HttpOnly: true,
		}
	}
	http.SetCookie(w, cookie)
	return nil
}
