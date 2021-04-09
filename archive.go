package archive

import (
	"bytes"
	"errors"
	readability "github.com/go-shiori/go-readability"
	"strings"
	"sync"
)



func Go(urlArr []string)( rs [] readability.Article , err error){

	if (len(urlArr)  > 100) {
		err = errors.New("数组长度最大为100")
		return  rs , err
	}

	var article readability.Article

	wg := sync.WaitGroup{}
	wg.Add(len(urlArr))

	chanArticleItem := make(chan readability.Article)

	for _ , value := range  urlArr {

		urlItem := value
		go func(string) {
			html := Crawl(urlItem)
			htmlReader := bytes.NewReader(html)
			article, err = readability.FromReader(htmlReader, urlItem)
			if err != nil {
				Error.Fatalf("failed to parse %s: %v\n", urlItem, err)
			}

			if article.Content == "" {
				Info.Println("静态页面获取失败，尝试动态页面获取", urlItem)
				// 如果静态页面的方式没获取到正文，说明是动态页面。
				html2, err := CrawlByRod(urlItem)
				if err != nil {
					Error.Printf("动态获取页面%s 出错，错误信息：%s\n", html2, err.Error())
				}
				htmlReader := strings.NewReader(html2)
				article, err = readability.FromReader(htmlReader, urlItem)
				if err != nil {
					Error.Fatalf("failed to parse %s: %v\n", urlItem, err)
				}
			}

			chanArticleItem <- article

		}(urlItem)
	}

	go func() {
		for i := range chanArticleItem{
			rs = append(rs , i )
			wg.Done()
		}
	}()

	wg.Wait()


	close(chanArticleItem)

	return rs , nil
}
