package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"html/template"
	"path"
	"runtime/debug"
)


const (
	LIST_DIR = 0x0001
	TEMPLATE_DIR = "./views"
	UPLOAD_DIR = "./uploads"
)

var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & LIST_DIR) ==0 {
			fi, err := os.Stat(file)
			if err != nil || fi.IsDir() {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHtml(w, "upload.html",nil )
		check(err)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename,
		http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)

	locals := make(map[string]interface{})
	images := []string{}
	for _,fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	err = renderHtml(w, "list.html", locals)
	check(err)
}

func renderHtml (w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	templates[tmpl].Execute(w, locals)
	return nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var imageId string
	var imagePath string
	imageId = r.FormValue("id")
	imagePath = UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	dat, err := ioutil.ReadFile(imagePath)
	if err == nil {
		stringType := http.DetectContentType(dat)
		w.Header().Set("Content-Type", stringType)
		http.ServeFile(w, r, imagePath)
	} else {
		log.Fatal("viewHandler:", err.Error())
	}
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func check (err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(),
					http.StatusInternalServerError)
				//或者输出自定义的50x错误页面
				//w.WriterHeader(http.StatusInternalServerError)
				//renderHtml(w, "error", e)
				//logging
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
//结合main()和uploadHandler()方法，针对HTTP GET 方式请求/upload路径，程序将会往http.ResponseWriter类型的实例对象w写入一段HTML文本，
//即输出一个HTML上传表单，如果使用浏览器访问这个地址，那么网页上将会是一个可以上传的文件的表单，接下来针对上传上来的图片在uploadHandler中进行操作
//如果是客户端发起的HTTP post请求，那么首先从表单提交过来的字段寻找名为image的文件域或者对应值，调用r.FormFile()方法会返回3个值
//各个值分别是multipart.File, *multipart.FileHeader和error
//如果上传的图片接收不成功，那么在示例程序中返回一个HTTP服务端的内部错误给客户端。
//如果上传的图片接收成功，那么将该图片的内容复制到一个临时文件里，如果临时文件创建失败，或者图片副本保存失败，都将触发服务器内部错误
//如果临时文件创建成功并且图片副本保存成功，即图片上传成功，就跳转到查看图片页面，此外，我们还定义了两个defer语句，无论图片上传成功或失败，
//当uploadHandler()方法执行结束时，都会先关闭临时文件句柄，继而关闭图片上传到服务器文件流的句柄

//当图片上传了，我们即可以在网页上查看这张图片，顺便确认图片是否真正上传到了服务端
//当图片上传成功后，我们必须有一个可查看的图片的网址，次例子显示图片上传成功后会跳转到/view?id=<imageId>这样的网址，因此我们程序要能够
//将对/view路径的访问映射到某个具体业务逻辑处理方法
/*
在view方法中，我们首先从客户端请求中对参数进行接值r.FromValue("id")即可得到客户端请求传递的图片的唯一ID，然后我们将图片ID结合之前保存图片用的
目录进行组装，即可得到文件在服务器上的存放路径，接着，调用http.ServeFile()方法将该路径下的文件从磁盘读取并作为服务端的返回信息输出给客户端，
同时也将HTTP响应头输出格式预设为image型，这是一种比较简单的方法，实际上应该严谨些，准备解析出文件的MimeType并将其为Content-Type进行输出，
具体参考http.DetectContentType()方法和mine包提供的相关方法,接下来将该方法注册到main中与访问路径形成映射关系

这样客户端访问/view路径并传递id参数时，即可直接以HTTP形式看到图片内容，在网页上，将会呈现出一张可视化的图片
理论上，只要是uploads/目录下有的图片都能够访问，但我们还是假设有意外情况，比如网址中传入了的图片ID在uploads/没有对应的文件，这时。这时我们的
方法就显得很脆弱了，不管是给出友好的错误提示还是返回404页面，都该对这种情况做相应处理，我们不妨先以最简单有效的方式对其进行处理，修改viewHandler()

列出所有已经上传图片
应该有个入口，可以看到所有已经上传的图片，对于所有累出的这些图片，我们可以选择进行查看或者删除操作，访问首页时列出所有上传的图片
由于我们将客户端上传的图片全部保存在工程的./uploads目录下，所以程序化总应该有个名叫listHandler()方法，用于在网页上列出该目录下存放的所有文件
暂时不考虑以缩略图的形式列出已上传图片，只需要列出可供访问的名称
listHandler()方法中可以看到。程序先从./uploads目录中遍历得到所有文件并赋值到fileInfoArr变量中，fileInfoArr是一个数组，其中的每一个
元素都是一个文件对象，然后，程序遍历fileInfoArr数组并从总得到图片的名称，用于后续的HTML片段中显示文件名和传入的参数内容，listHtml变量
用于在for循环中将图片名称一一串联起来生成一段HTML，最后调用io。WriteString()方法将这段HTML输出返回给客户端
关于HTML的方法都是用io.WriteString方法输出，在业务逻辑处理程序中混杂HTML可不是什么好事情，代码多了会导致程序不够清晰，而且改动程序里面的HTML
文本时，每次都要重新编译整个工程代码才能看到修改后的效果，应该将业务逻辑和编程分离开各自单独处理，这时候就需要前端技术了
Go提供html/template包，可以让我们将HTML从业务逻辑中抽离出来形成独立的模板文件，这样业务逻辑程序只负责处理业务和提供模板所需要的数据，
模板文件负责数据要表现的数据形式，然后模板解析器将这些数据以定义好的模板规则结合模板文件进行渲染，最终将渲染后的结构一并输出构成一个完整网页
加入templates进行HTML模板操作
*/
/*io.WriteString(w, "<html lang=\"en\"><head><meta charset=\"UTF-8\"></head>"+
			"<title>golang测试</title>"+
			"<form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\"/>"+
			"<input type=\"submit\" value=\"Upload\"/>"+
			"</form></html>")
*/
/*var listHtml string
	for _, fileInfo := range fileInfoArr {
		imgId := fileInfo.Name()
		listHtml += "<li><a href=\"/view?id="+imgId+"\">"+imgId+"</a></li>"
	}

	io.WriteString(w, "<html><ol>"+listHtml+"</ol></html>")
*/
//template.ParseFiles()函数将会读取指定模板的内容并且返回一个*template.Template值
//t.Execute()方法会根据模板语法来执行模板的渲染，并将渲染后的结果作为HTTP的返回数据输出
//在uploadHandler()方法和listHandler()方法中，均调用了template.ParseFiles()和t.Execute()这两个方法，根据DRY原则，我们可以将模板
//渲染代码分离出来，单独编写一个处理函数，以便其他业务逻辑处理函数都可以适应
/*
	t, err := template.ParseFiles("list.html")
	t.Execute(w, locals)


    t, err := template.ParseFiles("upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}


		t.Execute(w, nil)
*/
/*
当我们引入Go标准库中的html/template包，实现了业务逻辑层与表现层分离后，对模板渲染逻辑去重，编写并使用太平鸟共模板渲染方法，
这让业务逻辑看上去清晰简洁多，但这样每次两个方法去调用并重新渲染模板 ，很明显，效率很低，也比较浪费资源，对模板进行缓存处理可以解决着问题


对模板进行缓存，即指一次性预加在模板，我们可以在程序初始化的时候将所有模板一次性加载到程序中， 正好Go的包机制允许我们在init()函数中做这样的事情，
init()和main()函数之前执行
声明全局template用于存放模板内容
templates是一个map类型的复合机构，map的键是字符串类型，即模板名字，值是*template.Template类型，然后在init方法中一次性加载所有模板
在init方法中我们在template.ParseFiles()方法的外层强制使用templates.Must()进行封装，template.Must()确保了模板不能解析成功时，一定会
触发处理流程，之所以这样做，是因为倘若模板不能成功加载，程序能做的唯一有意义的事情就是退出
在range语句中，包含了我们希望加载的u和l.html两个模板，如果我们想加载更多模板，只需要往这个数组中加载更多元素既可以,当然最好的办法就是将所有
HTML模板文件统一放到一个文件夹中，然后对这个模板文件夹进行遍历和预加载，如果需要加载新的模板，只需在这个文件夹中新建模板即可这样的好处不用反复修改
即可重新编译程序，而且实现了业务层和表现层和业务层真正意义上的分离
*/
/*
	for _, tmpl := range []string{"upload", "list"} {
		t := template.Must(template.ParseFiles(tmpl + ".html"))
		templates[tmpl] = t
	}
*/
/*
func renderHtml (w http.ResponseWriter, tmpl string, locals map[string]interface{})(err error) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		return err
	}
	t.Execute(w, locals)
	return nil
}

包装check的方法来统一错误返回,但也带来一个问题，由于发生错误触发错误处理流程必然会导致程序停止运行，这种该发有点自欺欺人
换一种思维方式，尽管我们从树上写上能保证大多数错误都能得到相应的处理，但是墨菲和你过不去啊，如果程序正确的处理了99个错误，若有一个导致意外出现异常，
那程序终将停止，不管是什么错误只要触发处理流程那我们就有办法对其进行处理。
定义了一个名为safeHandler()的函数，该函数有一个参数并且返回一个值，偿还如的参数和返回值都是一个函数，且都是http.HandlerFunc了类型，
这种类型的函数有连个参数http.ResponseWriter和*http.Request函数规格同次例子的业务逻辑完全一致，事实上，我们正是要把业务逻辑处理函数作为
参数传入到这个方法中，这样任何一个错误处理流程向上回溯的时候，我们都能对其进行拦截处理，从而避免程序停止运行
在这个方法中巧妙的用了defer关键字配合recover方法中介panic的不断运行，该方法接收一个业务逻辑函数作为参数，同时调用这个业务逻辑函数，改业务逻辑执行完毕后
，该方法中的defer指定的匿名函数开始执行，倘若业务逻辑处理触发了panic则调用recover对其进行检测，若为一般性错误，则输出HTTP 50X出错信息并记录
日志，而程序将继续良好运行
*/