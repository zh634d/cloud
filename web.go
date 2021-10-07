package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func heathz(w http.ResponseWriter, r *http.Request)  {

	w.WriteHeader(200)
	fmt.Fprintf(w, "health check is ok", 200)

}
//将request的头部写入Response
func getheader(w http.ResponseWriter, r *http.Request) {

	for key := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", key, r.Header[key])
	}
   //输出部分示例
	w.Header().Set("version", os.Getenv("VERSION"))
	//w.WriteHeader(200)
	//w.Write([]byte("hello world\n"))
	//w.Header().Add("version",os.Getenv("VERSION"))
	w.Header().Add("name","cunsang")


	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "value: %s\n", r.Header["Accept-Encoding"])
	fmt.Fprintf(w, "value : %s\n", r.Header.Get("Accept-Encoding"))
	fmt.Fprintf(w, "value : %s\n", r.Header.Get("version"))
	fmt.Fprintf(w, "value : %s\n", r.Header.Get("name"))
	//将操作系统的VERSION写入response header
	//fmt.Fprintf(w, " value: %s\n", os.Getenv("VERSION "))


}
func loginfo()  {
	fmt.Println(log.Default())
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", getheader)
	http.HandleFunc("/healthz", heathz)
	server.ListenAndServe()
	loginfo()
}
