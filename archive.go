package archive

import (
	"fmt"
	"log"
	"regexp"

	"github.com/fatih/color"

	"github.com/MakeGolangGreat/archive-go/common"
	"github.com/MakeGolangGreat/archive-go/douban"
	"github.com/MakeGolangGreat/archive-go/weibo"
	"github.com/MakeGolangGreat/archive-go/weixin"
	"github.com/MakeGolangGreat/archive-go/zhihu"

	"github.com/MakeGolangGreat/telegraph-go"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Save 是一个备份函数，将链接内的文本抓取然后备份到Telegraph，然后返回一个Telegraph链接。
func Save(updateText string, token string, attachInfo *telegraph.NodeElement) (msg string, err error) {
	linkRegExp := regexp.MustCompile(`(http.*?)\s`)

	replyMessage := ""
	// 如果能匹配到某个链接
	// TODO 没有考虑到文章中有多个链接的可能，只是匹配了第一个
	if linkRegExp.MatchString(updateText) {
		// 拿到链接，但有可能是个错误的链接。

		fmt.Println("updateText", updateText)
		// Could be multi link inside the struct[][]string.
		matchURL := linkRegExp.FindAllSubmatch([]byte(updateText), -1)
		link := string(matchURL[0][1])

		page := &telegraph.Page{
			AccessToken: token,
			AuthorURL:   link,
			AuthorName:  projectName,
			AttachInfo:  attachInfo,
		}

		var err error
		fmt.Println(link)

		if zhihu.IsZhihuLink(link) {
			color.Green("监测到知乎链接")
			replyMessage, err = zhihu.Save(link, page)
		} else if douban.IsDoubanLink(link) {
			color.Green("监测到豆瓣链接")
			replyMessage, err = douban.Save(link, page)
		} else if weibo.IsWeiboLink(link) {
			color.Green("监测到微博链接")
			replyMessage, err = weibo.Save(link, page)
		} else if weixin.IsWeixinLink(link) {
			color.Green("监测到微信链接")
			replyMessage, err = weixin.Save(link, page)
		} else {
			color.Green("未适配该链接，走通用逻辑")
			replyMessage, err = common.Save(link, page)
		}

		if err != nil {
			return "", err
		}
	}
	return replyMessage, nil
}

// Text 是一个备份函数，将传递过来的文本备份到Telegraph，不管里面有没有链接，全部当成文本备份
// 然后返回一个Telegraph链接
func Text(updateText string, token string) (msg string, err error) {
	page := &telegraph.Page{
		AccessToken: token,
		AuthorURL:   projectLink,
		AuthorName:  projectName,
		Title:       "内容备份",
		Data:        updateText + projectDesc,
	}

	link, err := page.CreatePage()
	if err != nil {
		fmt.Println("保存文章失败", err)
		return "", err
	}

	return link, nil
}

// 检查此链接是否之前已经备份过，如果备份过，直接返回上次备份的链接
// 但不确定如何实现。关键在于如何保存每次的记录。本地数据库？那意味着将要长久地租一台服务器...
// 每次将保存记录保存在一个telegra.ph文章里？那么并发将是个问题，毕竟每次都要先读取telegra.ph链接来获取记录以及每次都要编辑telegra.ph文章。太频繁了。
func checkExist(link string) {
}
