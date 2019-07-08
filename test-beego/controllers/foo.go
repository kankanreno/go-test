package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type FooController struct {
	beego.Controller
}

//func (c *FooController) Prepare() {
//	fmt.Println("prepare...")
//
//	r := c.Ctx.Request
//	w := c.Ctx.ResponseWriter
//
//	if !cas.IsAuthenticated(r) {
//		//cas.RedirectToLogin(w, r)
//		w.Header().Set("Content-Type", "application/json")
//		str := fmt.Sprintf(`{"username":"%s","email":"kankan@pa.com"}`, cas.Username(r))
//		w.Write([]byte(str))
//		return
//	}
//
//	if r.URL.Path == "/logout" {
//		cas.RedirectToLogout(w, r)
//		return
//	}
//}

func (c *FooController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "foo.tpl"
}

func (c *FooController) GetJSON() {
	//c.Data["Website"] = "kankan.me"
	//c.Data["Email"] = "kankan@gmail.com"
	//c.TplName = "index.tpl"

	//c.Redirect("/login", 302)

	////// 模拟 wayne 返回
	mystruct := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		0,
		"this is message string",
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = &mystruct
	//c.Data["json"] = &map[string]interface{}{"code": 0, "message": "this is message string"}
	c.ServeJSON()
}
