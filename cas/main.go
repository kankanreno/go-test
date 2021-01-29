package main

import (
	"bytes"
	"fmt"
	"github.com/go-cas/cas"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"net/url"
)

const casURL = "http://localhost:8888/cas/"

type templateBinding struct {
	Username   string
	Attributes cas.UserAttributes
}

func main() {
	log.Info("=== Starting up")

	//// === BASE ===
	//m := http.NewServeMux()
	//m.HandleFunc("/foo", handlerFunc)
	//
	//server := &http.Server{
	//	Addr:    ":9999",
	//	Handler: m,
	//}
	//
	////if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	//if err := server.ListenAndServe(); err != nil {
	//	log.Infof("Error from HTTP Server: %v", err)
	//}

	//// === Wrap Sever ===
	//m := http.NewServeMux()
	//m.HandleFunc("/", handlerFunc)
	////m.Handle("/foo", middlewareLogger(http.HandlerFunc(fooHandlerFunc)))
	//
	////if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", m); err != nil {
	//if err := http.ListenAndServe(":9999", m); err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//// === SIMPLE ===
	//http.Handle("/", client.HandleFunc(handlerFunc))
	////if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", m); err != nil {
	//
	//if err := http.ListenAndServe(":9999", nil); err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//// === CAS ===
	//url, _ := url.Parse(casURL)
	//client := cas.NewClient(&cas.Options{URL: url})
	//
	//m := http.NewServeMux()
	//m.HandleFunc("/foo", fooHandlerFunc)
	//m.HandleFunc("/logout", fooHandlerFunc)
	////m.Handle("/foo", middlewareLogger(http.HandlerFunc(fooHandlerFunc)))
	////m.Handle("/logout", middlewareLogger(http.HandlerFunc(fooHandlerFunc)))
	//
	//server := &http.Server{
	//	Addr:    ":9999",
	//	Handler: client.Handle(middlewareLogger(m)),
	//}
	//
	////if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	//if err := server.ListenAndServe(); err != nil {
	//	log.Infof("Error from HTTP Server: %v", err)
	//}

	// === Middleware CAS ===
	m := http.NewServeMux()
	m.HandleFunc("/foo", fooHandlerFunc)
	m.HandleFunc("/bar", barHandlerFunc)

	server := &http.Server{
		Addr:    ":9999",
		Handler: middlewareCas(m),
	}

	//if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	if err := server.ListenAndServe(); err != nil {
		log.Infof("Error from HTTP Server: %v", err)
	}

	log.Info("=== Shutting down")
}

func middlewareCas(next http.Handler) http.Handler {
	//url, _ := url.Parse(casURL)
	//client := cas.NewClient(&cas.Options{URL: url})
	//return client.Handle(next)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("=== Executing middleware middlewareCas Start...")

		if r.URL.Path != "/noneedlogin" {
			url, _ := url.Parse(casURL)
			client := cas.NewClient(&cas.Options{URL: url})
			client.Handle(casHandler(next)).ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}

		log.Println("=== Executing middleware middlewareCas End...")
	})
}

func casHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("=== Executing middleware casHandler Start...")
		log.Infof("request: %+v", r)

		// 验证登录
		if !cas.IsAuthenticated(r) && r.URL.Path == "/currentuser" {
		//if !cas.IsAuthenticated(r) {
			// return HTML
			//cas.RedirectToLogin(w, r)
			//return

			//// return json
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Origin,Authorization,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Feedback-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Expose-Headers", "Feedback-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Feedback-Type")
			w.Header().Set("Feedback-Type", "application/json")
			w.WriteHeader(200)

			str := fmt.Sprintf(`{"code": -1, "message": "not logged in!"}`)
			log.Infof("=== Cas Middleware, not logged in! will return: %s", str)
			w.Write([]byte(str))
			return
		}

		// return to frontend, request come from cas server
		if r.URL.Path == "/redirect_to" {
			frontPath := r.URL.Query().Get("front_path")
			log.Infof("=== Cas Middleware, logged! will redirect to frontPath: %s", frontPath)
			http.Redirect(w, r, frontPath, http.StatusFound)
			return
		}

		if r.URL.Path == "/logout" {
			cas.RedirectToLogout(w, r)
			return
		}

		// 处理 /currentuser 请求，并将 token 的获取也统一到该请求中
		if r.URL.Path == "/currentuser" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Origin,Authorization,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Feedback-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Expose-Headers", "Feedback-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Feedback-Type")
			w.Header().Set("Feedback-Type", "application/json")
			w.WriteHeader(200)

			apiToken := "ttt"
			log.Infof("cas.Username(r): %s", cas.Username(r))
			jsonStr := fmt.Sprintf(`{"username": "%s"}`, cas.Username(r))
			str := fmt.Sprintf(`{"code": 0, "message": "%s", "data": %s}`, apiToken, jsonStr)
			w.Write([]byte(str))
			return
		}

		next.ServeHTTP(w, r)

		log.Println("=== Executing middleware casHandler End...")
	})
}

func fooHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Info("=== fooHandlerFunc...")
	w.Header().Add("Feedback-Type", "text/html")

	tmpl, err := template.New("index.html").Parse(index_html)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info(w, error_500, err)
		return
	}

	log.Info("cas.Username(r): ", cas.Username(r))
	binding := &templateBinding{
		Username:   cas.Username(r),
		Attributes: cas.Attributes(r),
	}

	html := new(bytes.Buffer)
	if err := tmpl.Execute(html, binding); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info(w, error_500, err)
		return
	}

	html.WriteTo(w)
}

const index_html = `<!DOCTYPE html>
<html>
  <head>
    <title>Welcome {{.Username}}</title>
  </head>
  <body>
    <h1>Welcome {{.Username}} <a href="/logout?url=http://localhost:9999">Logout</a></h1>
    <p>Your attributes are:</p>
    <ul>{{range $key, $values := .Attributes}}
      <li>{{$len := len $values}}{{$key}}:{{if gt $len 1}}
        <ul>{{range $values}}
          <li>{{.}}</li>{{end}}
        </ul>
      {{else}} {{index $values 0}}{{end}}</li>{{end}}
    </ul>
  </body>
</html>
`

const error_500 = `<!DOCTYPE html>
<html>
  <head>
    <title>Error 500</title>
  </head>
  <body>
    <h1>Error 500</h1>
    <p>%v</p>
  </body>
</html>
`

func barHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Info("=== barHandlerFunc...")
	w.Write([]byte("bar content"))
}
