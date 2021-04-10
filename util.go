package archive

// FixImage 有些网站，比如微信公众号，使用了动态加载图片的技术，这些页面的 img 标签的 href 属性不存在或者为空。真实的图片链接存储在 data-href（也不一定，需要程序推测） 这样的属性中。
// 本函数的作用就是把图片的真实链接放到 href 属性的值中。
func FixImage() {

}
