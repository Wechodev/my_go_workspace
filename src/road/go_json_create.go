package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
} 

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName:"shanghai_VPN", ServerIP:"127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName:"Beijing_VPN", ServerIP:"127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

//字段的tag是"_"，那么这个字段不会输出到JSON
//tag中带有自定义名称，那么这个子弟昂一名称会出现在JSON的字段名中，eg:serverName
//tag中带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
//如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
/*
1.JSON对象只支持string作为key，所有要编码一个map，那么必须是map[string]T这种类型（T是go语言中任意的类型）
2.Channel,complex和function是不能被编码成JSON的
3.嵌套的数据不能编码，不然会让json编码陷入死循环
4.指针在编码的时候会输出指针指向的内容，而空指针会输出null
*/