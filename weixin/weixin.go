package weixin

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	telegraph "github.com/MakeGolangGreat/telegraph-go"
)

func Save(url string, page *telegraph.Page) (link string, err error) {
	err2 := getContent(url, page)
	errHandler("获取内容失败", err2)
	link, err = page.CreatePage()
	return
}

func IsWeixinLink(url string) bool {
	reg := regexp.MustCompile(`http.*mp.weixin.qq.com.*`)
	return reg.MatchString(url)
}

func getContent(url string, data *telegraph.Page) error {
	spider := colly.NewCollector()
	extensions.RandomUserAgent(spider)

	spider.OnResponse(func(res *colly.Response) {
		dom, err := goquery.NewDocumentFromReader(bytes.NewReader((res.Body)))
		errHandler("初始化goquery失败", err)

		// 标题
		dom.Find("body #js_article #activity-name").Each(func(_ int, s *goquery.Selection) {
			data.Title = s.Text()
		})

		dom.Find("body #js_article #js_name").Each(func(_ int, s *goquery.Selection) {
			data.AuthorName = s.Text()
		})

		dom.Find("body #js_article #js_content").Each(func(_ int, s *goquery.Selection) {
			html, err := s.Html()
			errHandler("解析内容html失败", err)
			data.Data += html
		})

		// 文章时间在JS里，暂时不处理。
	})

	var err error
	// Set error handler
	spider.OnError(func(r *colly.Response, wrong error) {
		err = wrong
	})

	spider.Visit(url)

	return err
}

func errHandler(msg string, err error) {
	if err != nil {
		fmt.Printf("%s - %s\n", msg, err)
	}
}
