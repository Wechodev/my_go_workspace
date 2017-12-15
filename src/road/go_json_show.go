package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
} 

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
//首先查找tag含有Foo的可导出的struct字段(首字母大写)
//其次查找字段名是Foo的导出字段
//最后查找类似FOO或者FoO这样的除了首字母之前其他大小写不敏感的导出字段
