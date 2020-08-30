# archive-go

> 一个检测链接 -> 爬取数据 -> 备份到[Telegraph](http://telegra.ph/)的 Golang 库。
>
> 工作流程：给我一些字符串，我会检测其中的链接，然后启动爬虫程序，将链接中的信息抓下来然后上传到 Telegraph 平台上，然后将 Telegraph 的文章链接返回给你。
>
> 同时提供保存纯文本的接口，将传入的文本直接保存到 Telegraph 平台，然后返回一个可访问的链接。



### 接口

---

请移步[go doc]()



### 开始

---

1. 下载依赖

```go
go get -u github.com/MakeGolangGreat/archive-go
```

2. `test.go`

```go
package main

imoprt "github.com/MakeGolangGreat/archive-go"

func main(){
	link, err := archive.Save("<h1>html strings here</h1>", "...telegraph-token here...", "...attach info here...")
	if err != nil {
    fmt.Println("Save Article Failed: ", err)
  }else{
    fmt.Println(link)
  }
}
```


