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
	"math/rand"
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

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Index Handler")
	curTime := time.Now()
	aftTime := time.Now()
	clockOutTime := curTime.Add(time.Duration(9 * 60 * 60 * 1e9)).Format("2006-02-01 15:04:05 PM")
	clockInTime := aftTime.Format("2006-02-01 15:04:05 PM")
	write([]byte("Clock In:" + clockInTime + " Clock Out:" + clockOutTime + "\n"))

	b, err := ioutil.ReadFile("clockIn.dat")
	check(err)
	str := string(b)
	fmt.Fprintf(w, r.URL.Path[1:]+str)
}

func IconHandle(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Inside Icon Handler")
}

func RandomHandle(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Inside Random Handler")
	arrFood := []string{"冬菇鸡", "酱排骨", "鸡腿", "甜甜鸭", "酱油鸡", "水煮肉片",
		"酸菜鱼", "砂锅鸭", "虾", "可乐鸡", "金针菇鸡", "青瓜鸡", "焖排骨", "卤肉",
		"烧鸭腿", "猪蹄", "土豆片", "莴笋", "豆角", "娃娃菜", "茄子", "青菜", "汤饭7",
		"汤饭8", "粥", "牛腩面", "云吞", "蒸饺"}
	fmt.Fprintf(w, arrFood[rand.Intn(len(arrFood)-1)])
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
	http.HandleFunc("/", IndexHandle)
	http.HandleFunc("/favicon.ico", IconHandle)
	http.HandleFunc("/random", RandomHandle)

	//http.HandleFunc("/login", UserValidate)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
