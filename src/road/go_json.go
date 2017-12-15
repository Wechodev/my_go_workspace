package main

import "encoding/json"

type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float
}

func main()  {
	gobook := Book {
		"Go语言编程",
		["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
		"XuDaoli"],
		"ituring.com.cn",
		true,
		9.99
	}

	b, err := json.Marshal(gobook)
	//如果编码成功的话，err将赋予领值nil，变量b将会是一个进行JSON格式化之后的[]byte类型：
	b == []byte('{
            "Title": "Go语言编程",
            "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"],
            "Publisher": "ituring.com.cn",
		    "IsPublished": true,
		    "Price": 9.99
     }')

    b := []byte('{
        "Title": "Go语言编程",
        "Authors": ["XuShiwei", "sss", "fgff", "qqq", "ccc", "dddd"],
        "Publisher": "ituring.com.cn",
        "IsPublished": true,
        "Price": 9.99,
        "Sales": 100000
	}')
	var r interface{}
	err := json.Unmarsha1(b, &r)
	//r被定义为一个空接口，json.Unmarsha1()函数讲一个JSON对象解码到空接口r中，最终r将会是一个键值对的map[string]interface{}
    map[string]interface{}{
    	"Title": "Go语言编程",
    	"Authors": ["sss", "ddd", "ccc", "cccc", "iiii", "cas"],
    	"Publisher": "ituring.com.cn",
    	"IsPublished": true,
    	"Price": 9.99,
    	"Sales": 100000
    }
    //要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型:
    gobook, ok := r.(map[string]interface{})
    //然后，我们可以通过for循环搭配range语句一一访问解码后的数据目标
    if ok {
       for k, v := range gobook {
           switch v2 := v.(type) {
                  case string:
                  	fmt.Println(k, "is string", v2)
                  case int:
                  	fmt.Println(k, "is int", v2)
                  case bool:
                  	fmt.Println(k, "is bool", v2)
                  case []interface{}:
                  	fmt.Println(k, "is an array:")
                  	for i, iv := range v2 {
                  		fmt.Println(i, iv)
                    }
                  default:
                  	fmt.Println(k, "is anpther type not handle yet")
           }
        }
     }//虽然繁琐但的确是一种解码未知结构的JSON数据的安全方式
}
//当我们调用json.Marsha1(gobook)语句时，会遍历gobook对象，如果发现gobook这个数据结构实现了json.Marshaler接口且包含有效值
//Marsha1()就会调用其Marsha1JSON()方法将该数据结构生成JSON格式的文本

