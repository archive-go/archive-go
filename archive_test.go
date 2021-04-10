package archive

import (
	"fmt"
	"strings"
	"testing"
)

func TestGo(t *testing.T) {
	for _, url := range []string{"https://weibo.com/1533789260/K7N7B7vw1?ref=feedsdk&type=comment&i=1#_rnd1617874974706"} {
		article := Go(url)
		// fmt.Printf("Title   : %s\n", article.Title)
		// fmt.Printf("Author  : %s\n", article.Byline)
		// fmt.Printf("Length  : %d\n", article.Length)
		// fmt.Printf("Excerpt : %s\n", article.Excerpt)
		// fmt.Printf("SiteName: %s\n", article.SiteName)
		// fmt.Printf("Image   : %s\n", article.Image)
		// fmt.Printf("Favicon : %s\n", article.Favicon)
		fmt.Printf("Content : %s\n", article.Content)

		if strings.Trim(article.Content, " ") != "" {
			t.Log("测试成功！")
		} else {
			t.Error("测试失败")
		}
	}

}
