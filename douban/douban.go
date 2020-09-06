package douban

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	telegraph "github.com/MakeGolangGreat/telegraph-go"
)

// Save 通过处理知乎链接，保存到Telegraph，然后返回线上链接的入口函数。
func Save(url string, page *telegraph.Page) (link string, err error) {
	if isDoubanStatusLink(url) {
		// 备份单个豆瓣广告内容
		err2 := getSingleStatus(url, page)
		errHandler("获取豆瓣广告失败", err2)
		link, err = page.CreatePage()
	} else if isDoubanNoteLink(url) {
		// 备份单个专栏文章
		err2 := getSingleNote(url, page)
		errHandler("获取豆瓣日记失败", err2)
		link, err = page.CreatePage()
	}

	return
}

//IsDoubanLink 如果是豆瓣链接，返回true
func IsDoubanLink(url string) bool {
	reg := regexp.MustCompile(`http.*douban\.com.*`)
	return reg.MatchString(url)
}

// 豆瓣广播链接
func isDoubanStatusLink(url string) bool {
	reg := regexp.MustCompile(`http.*douban\.com\/people/.*/status\/\d+`)
	return reg.MatchString(url)
}

// 豆瓣日记链接
func isDoubanNoteLink(url string) bool {
	reg := regexp.MustCompile(`http.*douban\.com\/note\/\d+`)
	return reg.MatchString(url)
}

// 获取单独的豆瓣广播内容，爬虫解决静态页面
func getSingleStatus(url string, page *telegraph.Page) error {
	spider := colly.NewCollector()
	extensions.RandomUserAgent(spider)
	extensions.Referer(spider)

	spider.OnRequest(func(req *colly.Request) {
		fmt.Printf("fetching: %s\n", req.URL.String())
	})

	spider.OnResponse(func(res *colly.Response) {
		dom, err := goquery.NewDocumentFromReader(bytes.NewReader((res.Body)))
		errHandler("初始化goquery失败", err)

		// 标题
		dom.Find("#content h1").Each(func(_ int, s *goquery.Selection) {
			page.Title = s.Text()
		})

		// 豆瓣用户名
		dom.Find(".status-item .hd .lnk-people").Each(func(_ int, s *goquery.Selection) {
			page.AuthorName = s.Text()
		})

		// 广播内容
		dom.Find(".status-item .status-saying").Each(func(_ int, s *goquery.Selection) {
			html, err := s.Html()
			errHandler("解析内容html失败", err)
			page.Data += html
		})

		// 文章发布时间
		dom.Find(".status-item .hd .pubtime span").Each(func(_ int, s *goquery.Selection) {
			time := s.Text()
			// 在文章尾部增加发布时间
			page.Data += "<br/><blockquote>" + time + "</blockquote>"
		})
	})

	var err error
	// Set error handler
	spider.OnError(func(r *colly.Response, wrong error) {
		err = wrong
	})

	spider.Visit(url)

	return err
}

// 获取单独的豆瓣日记内容，爬虫解决静态页面
func getSingleNote(url string, page *telegraph.Page) error {
	spider := colly.NewCollector()
	extensions.RandomUserAgent(spider)
	extensions.Referer(spider)

	spider.OnRequest(func(req *colly.Request) {
		fmt.Printf("fetching: %s\n", req.URL.String())
	})

	spider.OnResponse(func(res *colly.Response) {
		dom, err := goquery.NewDocumentFromReader(bytes.NewReader((res.Body)))
		errHandler("初始化goquery失败", err)

		// 标题
		dom.Find(".note-container .note-header h1").Each(func(_ int, s *goquery.Selection) {
			page.Title = s.Text()
		})

		// 用户名
		dom.Find(".note-container .note-author").Each(func(_ int, s *goquery.Selection) {
			page.AuthorName = s.Text()
		})

		// 内容
		dom.Find(".note-container .note").Each(func(_ int, s *goquery.Selection) {
			html, err := s.Html()
			errHandler("解析内容html失败", err)
			page.Data += html
		})

		// 文章发布时间
		dom.Find(".note-container .pub-date").Each(func(_ int, s *goquery.Selection) {
			time := s.Text()
			// 在文章尾部增加发布时间
			page.Data += "<br/><blockquote>" + time + "</blockquote>"
		})
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
