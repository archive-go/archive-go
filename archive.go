package archive

import (
	"bytes"
	"strings"

	readability "github.com/go-shiori/go-readability"
)

func Go(url string) *readability.Article {
	var article readability.Article
	var err error
	html := Crawl(url)
	htmlReader := bytes.NewReader(html)
	article, err = readability.FromReader(htmlReader, url)
	if err != nil {
		Error.Fatalf("failed to parse %s: %v\n", url, err)
	}

	if article.Content == "" {
		Info.Println("静态页面获取失败，尝试动态页面获取", url)
		// 如果静态页面的方式没获取到正文，说明是动态页面。
		html2, err := CrawlByRod(url)
		if err != nil {
			Error.Printf("动态获取页面%s 出错，错误信息：%s\n", html2, err.Error())
		}
		htmlReader := strings.NewReader(html2)
		article, err = readability.FromReader(htmlReader, url)
		if err != nil {
			Error.Fatalf("failed to parse %s: %v\n", url, err)
		}
	}

	// fmt.Printf("URL     : %s\n", url)
	// fmt.Printf("Title   : %s\n", article.Title)
	// fmt.Printf("Author  : %s\n", article.Byline)
	// fmt.Printf("Length  : %d\n", article.Length)
	// fmt.Printf("Excerpt : %s\n", article.Excerpt)
	// fmt.Printf("SiteName: %s\n", article.SiteName)
	// fmt.Printf("Image   : %s\n", article.Image)
	// fmt.Printf("Favicon : %s\n", article.Favicon)
	// fmt.Printf("Content : %s\n", article.Content)

	return &article
}
