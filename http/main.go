package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"time"
)

type templateBinding struct {
	Username   string
	Attributes string
}

func main() {
	log.Info("=== Starting up")

	//// === BASE ===
	//mux := http.NewServeMux()
	//mux.HandleFunc("/foo", handlerFunc)
	//
	//server := &http.Server{
	//	Addr:    ":9999",
	//	Handler: mux,
	//}
	//
	////if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	//if err := server.ListenAndServe(); err != nil {
	//	log.Infof("Error from HTTP Server: %v", err)
	//}

	//// === WRAP SEVER ===
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", handlerFunc)
	////mux.Handle("/foo", middlewareLogger(http.HandlerFunc(fooHandlerFunc)))
	//
	////if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux); err != nil {
	//if err := http.ListenAndServe(":9999", mux); err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	// === SIMPLE ===
	http.HandleFunc("/", middlewareSimple(handlerFunc))
	//if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux); err != nil {

	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//// === CAS ===
	//url, _ := url.Parse(casURL)
	//client := cas.NewClient(&cas.Options{URL: url})
	//
	//mux := http.NewServeMux()
	//mux.HandleFunc("/foo", fooHandlerFunc)
	//mux.HandleFunc("/logout", fooHandlerFunc)
	////mux.Handle("/foo", middlewareLogger(http.HandlerFunc(fooHandlerFunc)))
	////mux.Handle("/logout", middlewareLogger(http.HandlerFunc(fooHandlerFunc)))
	//
	//server := &http.Server{
	//	Addr:    ":9999",
	//	Handler: client.Handle(middlewareLogger(mux)),
	//}
	//
	////if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	//if err := server.ListenAndServe(); err != nil {
	//	log.Infof("Error from HTTP Server: %v", err)
	//}

	//// === Middleware CAS ===
	//mux := http.NewServeMux()
	//mux.HandleFunc("/foo", fooHandlerFunc)
	//mux.HandleFunc("/bar", barHandlerFunc)
	//
	//server := &http.Server{
	//	Addr:    ":9999",
	//	Handler: middlewareCas(mux),
	//}
	//
	////if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
	//if err := server.ListenAndServe(); err != nil {
	//	log.Infof("Error from HTTP Server: %v", err)
	//}
	//
	//log.Info("=== Shutting down")
}

func now() string {
	return time.Now().Format(time.Stamp) + " "
}

func middlewareSimple(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(now() + "before")
		defer fmt.Println(now() + "after")
		f(w, r)
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("h4")
}

func middlewareCas(next http.Handler) http.Handler {
	//url, _ := url.Parse(casURL)
	//client := cas.NewClient(&cas.Options{URL: url})
	//return client.Handle(next)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("=== Executing middleware middlewareCas Start...")

		//if r.URL.Path != "/noneedlogin" {
		//	url, _ := url.Parse(casURL)
		//	client := cas.NewClient(&cas.Options{URL: url})
		//	client.Handle(casHandler(next)).ServeHTTP(w, r)
		//} else {
		//	next.ServeHTTP(w, r)
		//}

		log.Println("=== Executing middleware middlewareCas End...")
	})
}

func fooHandlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Info("=== fooHandlerFunc...")
	w.Header().Add("Reply-Type", "text/html")

	tmpl, err := template.New("index.html").Parse(index_html)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info(w, error_500, err)
		return
	}

	binding := &templateBinding{
		Username:   "kankan",
		Attributes: "123",
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
