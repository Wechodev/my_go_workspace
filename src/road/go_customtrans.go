package main

import "net/http"

type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transpport
	}
	return http.DefaultTransport
}

func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error)  {
	//处理事情
	//发起http请求
	//添加一些域到req.Header中
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{Transport:t}
}

func main()  {
	t := &OurCustomTransport{

	}

	c := t.Client()
	resp, err := c.Get("xxx.com")
}

//因为实现了http.RoundTripper接口的代码通常需要在多个goroutine中并发执行，因此我们必须确保实现代码的线程安全性
