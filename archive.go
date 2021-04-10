package archive

import (
	"io"

	readability "github.com/go-shiori/go-readability"
)

func Go(url string) (article readability.Article, err error) {
	var htmlReader io.Reader

	htmlReader, err = Crawl(url)
	if err != nil {
		Error.Printf("爬取页面出错：%s , 错误信息：%s", url, err.Error())
	} else {
		// 只有在静态爬虫没出错的情况下才尝试解析HTML
		article, err = readability.FromReader(htmlReader, url)
		if err != nil {
			Error.Printf("failed to parse %s : %v\n", url, err)
		}
	}

	if article.Content == "" {
		Info.Println("静态页面获取失败，尝试动态页面获取", url)
		// 如果静态页面的方式没获取到正文，说明是动态页面。
		htmlReader, err = CrawlByRod(url)
		if err != nil {
			Error.Printf("动态获取页面%s 出错，错误信息：%s\n", url, err.Error())
			return
		}
		article, err = readability.FromReader(htmlReader, url)
		if err != nil {
			Error.Printf("failed to parse %s : %v\n", url, err)
		}
	}

	return
}
