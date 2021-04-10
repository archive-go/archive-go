package archive

import (
	"context"
	"time"

	"github.com/go-rod/rod"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var browser = rod.New().MustConnect()

// Crawl 假设目标网页为静态资源，获取网页HTML
func Crawl(url string) (html []byte) {
	spider := colly.NewCollector()
	extensions.RandomUserAgent(spider)
	extensions.Referer(spider)

	spider.OnResponse(func(res *colly.Response) {
		html = res.Body
	})

	// Set error handler
	spider.OnError(func(r *colly.Response, err error) {
		Error.Fatalf("Colly 爬虫出错，URL： %s， 错误信息：%s\n", url, err.Error())
	})

	spider.Visit(url)

	return
}

func CrawlByRod(url string) (html string, err error) {
	page := browser.MustPage(url)
	ctx, cancel := context.WithCancel(context.Background())
	pageWithCancel := page.Context(ctx)
	go func() {
		time.Sleep(10 * time.Second)
		Warning.Printf("Rod 爬虫超时，URL：%s\n", url)
		cancel()
	}()
	html, err = pageWithCancel.MustWaitLoad().HTML()
	if err != nil {
		Warning.Printf("Rod 爬虫出错，URL：%s，错误信息：%s\n", url, err.Error())
	}
	return
}
