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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(s []byte) {
	fd, err := os.OpenFile("clockIn.dat", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	check(err)
	_, err = fd.Write(s)
	check(err)
}

func Index(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now()
	aftTime := time.Now()
	clockOutTime := curTime.Add(time.Duration(8 * 60 * 60 * 1e9)).Format("2006-02-01 15:04:05 PM")
	clockInTime := aftTime.Format("2006-02-01 15:04:05 PM")
	write([]byte("Clock In:" + clockInTime + "Clock Out:" + clockOutTime + "\n"))

	fmt.Println("Inside Index handler")
	b, err := ioutil.ReadFile("clockIn.dat")
	check(err)
	str := string(b)
	fmt.Fprintf(w, r.URL.Path[1:]+str)
}

func Icon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Icon handler")
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
	http.HandleFunc("/favicon.ico", Icon)

	//http.HandleFunc("/login", UserValidate)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
