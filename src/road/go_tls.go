package main

import (
	"fmt"
	"net/http"
)

/*const SERVER_PORT = 8081
const SERVER_DOMIAN = "localhost"
const RESPONSE_TEMPLATE = "hello"*/

/*func rootHandler(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}*/

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMIAN, SERVER_PORT), rootHandler)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "cert.pem", "key.pem", nil)
}




//运行该服务器后，我们可以在浏览器中放问8080端口查看效果
/*
可以看到我们使用了TLS这个方法，这表明它执行在TLS层上的HTTP协议，我们并不需要支持HTTPS，只需要把方法的TLS去掉即可
*/


