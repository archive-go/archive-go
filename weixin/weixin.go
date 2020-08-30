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
	// extensions.Referer(spider)

	// spider.OnRequest(func(req *colly.Request) {
	// 	req.Headers.Set("Cookie", "pac_uid=0_c72e541fad774; iip=0; pgv_pvid=9044765905; RK=0cDhfeMnRL; ptcz=cf8b83f8a151ce68832c9e5f62e64fd725d94ee65732d84f490296bda3095881; pgv_pvi=4004080640; rewardsn=; wxtokenkey=777")
	// 	req.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// 	req.Headers.Set("Accept-Encoding", "gzip, deflate, br")
	// 	req.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	// 	req.Headers.Set("Cache-Control", "no-cache")
	// 	req.Headers.Set("Connection", "keep-alive")
	// 	req.Headers.Set("DNT", "1")
	// 	req.Headers.Set("Host", "mp.weixin.qq.com")
	// 	req.Headers.Set("Pragma", "no-cache")
	// 	req.Headers.Set("Sec-Fetch-Dest", "document")
	// 	req.Headers.Set("Sec-Fetch-Mode", "navigate")
	// 	req.Headers.Set("Sec-Fetch-Site", "none")
	// 	req.Headers.Set("Sec-Fetch-User", "71")
	// 	req.Headers.Set("Upgrade-Insecure-Requests", "1")
	// 	req.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_16_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36")
	// })

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
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", "\nError:", wrong)
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
