package main

import (
	"bytes"
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
		Handler: middlewareCas(middlewareLogger(m)),
	}

	//if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	if err := server.ListenAndServe(); err != nil {
		log.Infof("Error from HTTP Server: %v", err)
	}

	log.Info("=== Shutting down")
}

// TODO: 改进：针对特定路由设置 cas client
func middlewareCas(next http.Handler) http.Handler {
	url, _ := url.Parse(casURL)
	client := cas.NewClient(&cas.Options{URL: url})
	return client.Handle(next)

	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Info("Executing middleware middlewareCas Start...")
	//
	//	if r.URL.Path == "/foo" {
	//		url, _ := url.Parse(casURL)
	//		client := cas.NewClient(&cas.Options{URL: url})
	//		client.Handle(next).ServeHTTP(w, r)
	//	} else {
	//		next.ServeHTTP(w, r)
	//	}
	//
	//	log.Println("Executing middleware middlewareCas End...")
	//})
}

// TODO: 路由判断
func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Executing middleware middlewareLogger Start...")

		if !cas.IsAuthenticated(r) && r.URL.Path != "/bar" {
			cas.RedirectToLogin(w, r)
			return
		}

		if r.URL.Path == "/logout" {
			cas.RedirectToLogout(w, r)
			return
		}

		next.ServeHTTP(w, r)

		log.Println("Executing middleware middlewareLogger End...")
	})
}

func fooHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Info("=== fooHandlerFunc...")
	w.Header().Add("Content-Type", "text/html")

	tmpl, err := template.New("index.html").Parse(index_html)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info(w, error_500, err)
		return
	}

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
    <h1>Welcome {{.Username}} <a href="/logout">Logout</a></h1>
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
