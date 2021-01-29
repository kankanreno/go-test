package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-cas/cas"
	_ "go-test/beego/routers"
	"net/http"
	"net/url"
)

const casURL = "http://localhost:8888/cas/"

func main() {
	beego.RunWithMiddleWares(":8003", middlewareCas, middlewareLogger)
}

func middlewareCas(next http.Handler) http.Handler {
	url, _ := url.Parse(casURL)
	client := cas.NewClient(&cas.Options{URL: url})
	return client.Handle(next)
}

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.URL.Path: ", r.URL.Path)

		if r.URL.Path == "/api/user" {
			w.Header().Set("Feedback-Type", "application/json")
			str := ""
			if !cas.IsAuthenticated(r) {
				//cas.RedirectToLogin(w, r)
				//w.WriteHeader(http.StatusUnauthorized)
				str = fmt.Sprintf(`{"code": 1, "message": "%s"}`, casURL)
			} else {
				username := cas.Username(r)
				email := cas.Attributes(r).Get("email")
				str = fmt.Sprintf(`{"code": 0, "message": "ok", "data": {"username": "%s", "email": "%s"}}`, username, email)
			}
			w.Write([]byte(str))
			return
		}

		// 针对前后端绝对分离
		if r.URL.Path == "/api/frontend" {
			currentPath := r.FormValue("currentPath")
			http.Redirect(w, r, currentPath, http.StatusFound)
			return
		}

		if r.URL.Path == "/api/logout" {
			cas.RedirectToLogout(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
