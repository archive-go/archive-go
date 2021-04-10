package archive

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var browser *rod.Browser

func init() {
	u := launcher.New().
		Set("--headless").
		Set("--disable-gpu").
		Set("--disable-dev-shm-usage").
		Set("--disable-setuid-sandbox").
		Set("--no-first-run").
		Set("--no-sandbox").
		Set("--no-zygote").
		Set("--single-process").
		MustLaunch()
	fmt.Println("_________________init__________________")
	browser = rod.New().ControlURL(u).MustConnect()
}

// Crawl 假设目标网页为静态资源，获取网页HTML
func Crawl(url string) (htmlReader *bytes.Reader, err error) {
	spider := colly.NewCollector()
	extensions.RandomUserAgent(spider)
	extensions.Referer(spider)

	spider.OnResponse(func(res *colly.Response) {
		htmlReader = bytes.NewReader(res.Body)
	})

	spider.OnError(func(r *colly.Response, _err error) {
		Error.Printf("Colly 爬虫出错，URL： %s ， 错误信息：%s\n", url, _err.Error())
		err = _err
	})

	spider.Visit(url)

	return
}

func CrawlByRod(url string) (htmlReader *strings.Reader, err error) {
	var page *rod.Page
	page, err = browser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	pageWithCancel := page.Context(ctx)

	go func() {
		time.Sleep(20 * time.Second)
		Warning.Printf("Rod 爬虫超时，URL：%s\n", url)
		cancel()
	}()

	//如果是微博 ， 处理跳转问题
	if rs := strings.Contains(url, "weibo"); rs {
		err = pageWithCancel.Wait(nil, "document.querySelectorAll('div').length > 10 ", nil)
		if err != nil {
			return
		}
	}
	err = pageWithCancel.WaitLoad()
	if err != nil {
		Error.Println("WaitLoad 出错，页面URL：", url)
		return
	}

	var html string
	html, err = pageWithCancel.HTML()
	if err != nil {
		Warning.Printf("Rod 爬虫出错，URL：%s ，错误信息：%s\n", url, err.Error())
		return
	}
	htmlReader = strings.NewReader(html)

	return
}
