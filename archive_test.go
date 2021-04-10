package archive

import (
	"strings"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestGo(t *testing.T) {
	// urls := []string{"https://www.zhihu.com/question/22199390/answer/990830919", "https://www.zhihu.com/question/449620519/answer/1806307276", "https://www.zhihu.com/question/453279708/answer/1825027518", "https://www.douban.com/group/topic/220042478/", "https://www.douban.com/group/topic/219888045/", "https://www.douban.com/group/topic/219565261/", "https://dig.chouti.com/link/30634040", "https://mp.weixin.qq.com/s/mfuuFr1s3qC1o3oRtDt1yQ", "https://mp.weixin.qq.com/s/HgzKdXUm4fNI6uIinbuVEw", "https://movie.douban.com/subject/6863983/", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment"}
	urls := []string{"https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment"}
	urls = append(urls, urls...)
	urls = urls[:len(urls)-1]
	t.Log("一共有多少条数据：", len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			article, _ := Go(link)
			// fmt.Printf("URL   : %s\n", link)
			// fmt.Printf("Title   : %s\n", article.Title)
			// fmt.Printf("Author  : %s\n", article.Byline)
			t.Logf("Length  : %d\n", article.Length)
			// fmt.Printf("Excerpt : %s\n", article.Excerpt)
			// fmt.Printf("SiteName: %s\n", article.SiteName)
			// fmt.Printf("Image   : %s\n", article.Image)
			// fmt.Printf("Favicon : %s\n", article.Favicon)
			// t.Logf("Content : %s\n", article.TextContent)
			// fmt.Println("_______________")

			if strings.Trim(article.Content, " ") != "" {
				t.Log("测试成功！")
			} else {
				t.Error("测试失败")
			}
		}(url)
	}

	wg.Wait()
}
