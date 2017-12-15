package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Hello, world")
}

func main()  {
	http.HandleFunc("/hello", helloHandler)
	//该方法用于分发请求，即针对某一路径将其映射到指定的业务逻辑方法中，即是URL路由
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

//在hello.go中，http.HandleFunc()方法接受两个参数，第一个参数是HTTP请求的目标路径"/hello",该值可以是字符串
//也可是字符串形式的正则表达式，第二个参数指定具体的回调方法，比如helloHandler,当我们程序运行起来后访问127.0.0.1/8081/hello
//程序就会去调用helloHandler()方法中的业务逻辑程序

/*
helloHandler()方法是http.HandlerFunc类型的实例，并传入http.ResponseWriter和http.Request作为其必要的两个参数，
http.ResponseWriter类型的对象用于包装处理HTTP服务端的响应信息，我们将字符串"Hello, world!"写入类型为http.ResponseWriter
的w实例中，即可将该字符串数据发送到HTTP客户端，第二个参数r *http.Request表示的是此次HTTP请求的一个数据结构体，即代表一个客户端，
不过我们尚未用到它

我们还看到，在main()方法调用了http.ListenAndServe()，该方法用于在示例中监听8081端口，接受，并调用内部程序来处理连接到此
端口的请求，如果端口监听失败，会调用log.Fatal()方法输出异常出错信息
*/
