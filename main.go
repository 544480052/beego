package beego

import (
	"net/http"//http请求包
	"fmt"
)

func main()  {
	//将9090端口的请求转派到home函数
	http.HandleFunc("/",home)
	http.ListenAndServe(":9090",nil)
}

func home(w http.ResponseWriter,r *http.Request)  {
	//输出给客户端信息
	fmt.Fprintf(w,"hello world!!!!")
}