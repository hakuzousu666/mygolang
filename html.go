package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>HanToolsBlog</title>
</head>
	这个页面依然很丑对吧？<br>
	没事，达到效果就行，对吧？<br>
	这个页面就是作者的开发日记<br>
	第 1 篇《扩展》2021/4/3：<br>
	这时的HanTools已经有了两个功能<br>
	用户一定会感觉很少<br>
	所以，我添加了一个功能<br>
	这个功能就是扩展！<br>
	可这个扩展有一个问题<br>
	就是只能添加Python文件<br>
	所以为什么呢？<br>
	因为我太懒了...<br>
	哔~<br>
	当然，今天我还不要脸的把这个页面发布了出去...<br>
	额...我觉得以这种形式可以去表达我写代码的时候的心情<br>
	用户可以更好的理解开发者<br>
	像：<br>
	你更新怎么怎么慢？<br>
	你就更新了怎么点东西？<br>
	不会吧？不会吧？不会怎么垃圾的东西还会有人用吧？<br>
	......<br>
	这种问题也会逐渐减少<br>
	今天就先说到这吧，明天再见！
<body>
</body>
</html>

`
