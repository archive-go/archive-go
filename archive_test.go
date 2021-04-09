package archive

import (
	"fmt"
	"strings"
	"testing"
)

func TestGo(t *testing.T) {

	//url := "https://www.jianshu.com/p/2ec181648d9a"
	//url := "https://weibo.com/5466550668/Ka4M2wzRk"
	url := "https://weibo.com/u/1557605344"
	for _, url := range []string{url} {
		article := Go(url)
		fmt.Printf("Title   : %s\n", article.Title)
		fmt.Printf("Author  : %s\n", article.Byline)
		fmt.Printf("Length  : %d\n", article.Length)
		fmt.Printf("Excerpt : %s\n", article.Excerpt)
		fmt.Printf("SiteName: %s\n", article.SiteName)
		fmt.Printf("Image   : %s\n", article.Image)
		fmt.Printf("Favicon : %s\n", article.Favicon)
		fmt.Printf("Content : %s\n", article.Content)

		if strings.Trim(article.Content, " ") != "" {
			t.Log("测试成功！")
		} else {
			t.Error("测试失败")
		}
	}

}
