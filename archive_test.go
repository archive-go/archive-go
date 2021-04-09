package archive

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {


	urls := []string{"https://www.zhihu.com/question/22199390/answer/990830919", "https://www.zhihu.com/question/449620519/answer/1806307276", "https://www.zhihu.com/question/453279708/answer/1825027518", "https://www.douban.com/group/topic/220042478/", "https://www.douban.com/group/topic/219888045/", "https://www.douban.com/group/topic/219565261/", "https://dig.chouti.com/link/30634040", "https://mp.weixin.qq.com/s/mfuuFr1s3qC1o3oRtDt1yQ", "https://mp.weixin.qq.com/s/HgzKdXUm4fNI6uIinbuVEw", "https://movie.douban.com/subject/6863983/", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment", "https://weibo.com/5466550668/Ka4M2wzRk", "https://m.weibo.cn/detail/4618414788445382", "https://weibo.com/6003416687/Ka69owNyW?type=comment"}
	rs , err :=  Go(urls)
	if err != nil {
		fmt.Println(err)
	}
	for _ , v := range rs  {
		fmt.Println(v.Title)
	}
}

