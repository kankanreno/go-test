package routers

import (
	"github.com/astaxie/beego"
	"test-beego/controllers"
)

const URL_LOGIN = "http://localhost:8888/cas/"

func init() {
	//url, _ := url.Parse(URL_LOGIN)
	//client := cas.NewClient(&cas.Options{URL: url})
	//beego.Handler("/", client.HandleFunc(logger), true)

	//beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
	//	fmt.Println("filter...")
	//
	//	//_, ok := ctx.Input.Session("uid").(int)
	//	//if !ok && ctx.Request.RequestURI != "/login" {
	//	//	ctx.Redirect(302, "/login")
	//	//}
	//
	//	r := ctx.Request
	//	w := ctx.ResponseWriter
	//
	//	//fmt.Printf("request: %+v\n", r)
	//
	//	if !cas.IsAuthenticated(r) {
	//		cas.RedirectToLogin(w, r)
	//		return
	//	}
	//
	//	if r.URL.Path == "/logout" {
	//		cas.RedirectToLogout(w, r)
	//		return
	//	}
	//})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/foo", &controllers.FooController{})
	beego.Router("/foojson", &controllers.FooController{}, "get,post:GetJSON")
}

//func logger(w http.ResponseWriter, r *http.Request) {
//	fmt.Print("request r: %+v", r)
//	if !cas.IsAuthenticated(r) {
//		cas.RedirectToLogin(w, r)
//		return
//	}
//
//	if r.URL.Path == "/logout" {
//		cas.RedirectToLogout(w, r)
//		return
//	}
//}
