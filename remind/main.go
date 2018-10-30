//package remind
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func HelloServer(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("Inside HelloServer handler")
//	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
//}
//
//type myMux struct {}
//
//func (p *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path == "/" {
//		HelloServer(w, r)
//		return
//	}
//
//	http.NotFound(w, r)
//	return
//}
//
//
//func remind() {
//	mux := &myMux{}
//
//	err := http.ListenAndServe("localhost:8081", mux)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err.Error())
//	}
//}
package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+r.URL.Path[1:])
}

//func UserValidate(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println("method", r.Method)
//	for v, k := range r.Form {
//		fmt.Println("value", v)
//		fmt.Println("key", k)
//	}
//}

func main() {
	http.HandleFunc("/", Index)
	//http.HandleFunc("/login", UserValidate)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
