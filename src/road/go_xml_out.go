package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)
}
//如果v是array或者slice,那么输出每一个元素，类似value
//如果v是指针，那么会Marsha1指针指向的内容，如果指针为空，什么都不输出
//如果v是interface那么久处理interface所包含的数据
//如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息
/*
生成的XML文件中的element的名字：
如果v是struct，XMLName的tag中定义的名称
类型为xml.Name的名叫XMLName的字段的值
通过struct中字段的tag来获取
通过struct的字段名用来获取
Marshall的类型名称

设置struct中字段的tag信息以控制最终的xml文件的生成
XMLName 不会被输出
tag中含有"-"的字段不会被输出
tag中含有"name,attr"，会以name作为属性名，字段值为输出为这个XML元素的属性，eg:version字段所描述
tag中含有",attr",会以这个struct的字段名作为属性名输出为XML元素的属性，只是这个name是默认字段名
tag中含有".innerxml",将会被原样输出，而不会进行常规的编码过程
tag中含有".comment",将被当做xml注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串
tag中含有",omitempty"如果该字段值为空那么该字段就不会被输出到XML，包括:false,0,nil,空指针，nil接口，任何长度为0的array，slice，string
tag中含有"a>b>c"那么就会循环输出这三个元素a包含b包含c
*/
