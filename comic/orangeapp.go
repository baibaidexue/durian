package comic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"durian/dao"
	"durian/helpers"
	"durian/webbody"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html"
)

const (
	OrangeAppCCURLBase = "https://orangeapp.cc/"
)

type ComicShort struct {
	Title       string
	Description string
}

func isHeadTitleNode(node *html.Node) bool {
	if node.Data == "title" {
		return true
	}
	return false
}

func isHeadDescNode(node *html.Node) bool {
	var descNode bool
	if node.Data == "meta" {

		for _, at := range node.Attr {
			if at.Key == "name" && at.Val == "description" {
				descNode = true
			}
		}
	}
	return descNode
}

func comicShortFind(doc *goquery.Document) (info ComicShort) {
	head := doc.Find("head")
	for _, h := range head.Nodes {
		for node := h.FirstChild; node != nil; node = node.NextSibling {
			fmt.Printf("headNode: %+v\n", node)
			if isHeadTitleNode(node) {
				if node.FirstChild != nil {
					info.Title = node.FirstChild.Data
				}
			} else if isHeadDescNode(node) {
				for _, v := range node.Attr {
					if v.Key == "content" {
						info.Description = v.Val
					}
				}
			}
		}
	}

	return
}

func nodeContainClass(node *html.Node, className string) bool {
	for _, v := range node.Attr {
		if v.Key == "class" {
			if v.Val == className {
				return true
			}
		}

	}
	return false
}

func nodeHasAttr(node *html.Node, attr string) (bool, string) {
	for _, v := range node.Attr {
		if v.Key == attr {
			return true, v.Val
		}
	}
	return false, ""
}

func whatMetaNode(node *html.Node) string {
	if nodeContainClass(node, "p-t-5 p-b-5 read-block") {
		return "read-block"
	}

	var count int
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		count++
	}

	// total 10 with tag block
	if count >= 8 {
		return "tag-block"
	}

	return "unhandled"
}

func orangeAppCrackName(p *goquery.Selection) (bookname string) {
	head := p.Find(".panel-heading")

	for _, b := range head.Nodes {
		for node := b.FirstChild; node != nil; node = node.NextSibling {
			if ok, val := nodeHasAttr(node, "itemprop"); ok {
				// fmt.Println("panel heading child alike book name node val:", val)
				if val == "name" {
					// fmt.Println("Val name checking")
					for nodeI := node.FirstChild; nodeI != nil; nodeI = nodeI.NextSibling {
						if nodeI.Data == "h1" {
							// fmt.Println("h1 finded")
							// fmt.Printf("h1 node:\n%+v\n", nodeI)
							for d := nodeI.FirstChild; d != nil; d = d.NextSibling {
								// fmt.Printf("h1 child nodes:\n%+v\n", d)
								bookname = d.Data
							}
						}

					}
				}

			}
		}

	}
	return
}

func OrangeAppCrach(doc *goquery.Document) (meta *dao.ComicMeta, err error) {
	meta = &dao.ComicMeta{}
	panel := doc.Find(".col-md-12")
	if len(panel.Nodes) == 0 {
		return nil, errors.New("Nothing found with comic url with first col md 12")
	}

	// name
	bookName := orangeAppCrackName(panel)
	meta.Name = bookName

	// poster image
	postserUrl := comicPosterFind(panel)
	meta.PosterURL = postserUrl.PosterUrl
	if meta.PosterURL == "" {
		return nil, errors.New("Nothing found with comic url on poster")
	}

	// tag & download & other info
	OrangeAppMetaFind(panel, meta)

	return meta, nil
}

func dealReadBlock(b *html.Node, meta *dao.ComicMeta) {
	for node := b.FirstChild; node != nil; node = node.NextSibling {
		if ok, val := nodeHasAttr(node, "style"); ok {
			if val == "padding: 5px;" {
				if ok, href := nodeHasAttr(node, "href"); ok {
					pa, _ := url.JoinPath(OrangeAppCCURLBase, href)
					meta.Ep = append(meta.Ep, dao.ComicEPInfo{
						DownloadUrl: pa,
						FilePath:    "",
						FileName:    "单本",
						Size:        0,
						Healthy:     0,
					})

				}
			}
			continue
		}
		if nodeContainClass(node, "btn-group") {
			for n := node.FirstChild; n != nil; n = n.NextSibling {
				if nodeContainClass(n, "dropdown-menu") {
					for li := n.FirstChild; li != nil; li = li.NextSibling {
						for lic := li.FirstChild; lic != nil; lic = lic.NextSibling {
							for licd := lic.FirstChild; licd != nil; licd = licd.NextSibling {
								_, href := nodeHasAttr(lic, "href")
								if href == "" {
									continue
								}
								downLink, _ := url.JoinPath(OrangeAppCCURLBase, href)
								name := strings.Trim(licd.Data, "\n")
								meta.Ep = append(meta.Ep, dao.ComicEPInfo{
									DownloadUrl: downLink,
									FilePath:    "",
									FileName:    name,
									Size:        0,
									Healthy:     0,
								})
							}
						}
					}
				}
			}
		}
	}
}

func OrangeAppMetaFind(p *goquery.Selection, meta *dao.ComicMeta) {
	panel := p.Find(".col-lg-7")

	for _, b := range panel.Nodes {
		for node := b.FirstChild; node != nil; node = node.NextSibling {

			switch whatMetaNode(node) {
			case "tag-block":

				dealTagInfo(node, meta)

			case "read-block":
				dealReadBlock(node, meta)
			}

		}
	}

}

type tagInfo struct {
	Name string
	Link string
}

func tagInfoInnerGet(in *html.Node) (da []tagInfo) {
	for d := in.FirstChild; d != nil; d = d.NextSibling {
		var href, val string
		if ok, v := nodeHasAttr(d, "href"); ok {
			href = v
		}

		if d.Data == "a" {
			valnode := d.FirstChild
			if valnode == nil {
				fmt.Println("a had a nil node")
			} else {
				val = valnode.Data
			}
		}
		if val != "" {
			da = append(da, tagInfo{Name: val, Link: href})
		}
	}
	return
}

func formatTime(dateStr string) time.Time {

	layout := "2006-01-02" // 指定日期字符串的格式
	t, _ := time.Parse(layout, dateStr)
	return t
}

func dealTagInfo(b *html.Node, meta *dao.ComicMeta) {
	for node := b.FirstChild; node != nil; node = node.NextSibling {
		if nodeContainClass(node, "p-t-5 p-b-5") {
			if node.Type == html.TextNode {
				fmt.Printf("text node: %+v\n", node)
			}
			if node.Type == html.ElementNode {
				// fmt.Printf("element node: %+v\n", node)
				for in := node.FirstChild; in != nil; in = in.NextSibling {
					// fmt.Printf("child data:\n%+v\n", in)
					// fmt.Println(in.Data)
					if strings.HasPrefix(in.Data, "\n禁漫车：") {
						meta.JMID = strings.TrimPrefix(in.Data, "\n禁漫车：")
					}
					if strings.HasPrefix(in.Data, "\n叙述：") {
						meta.Overview = strings.TrimPrefix(in.Data, "\n叙述：")
					}
					if strings.HasPrefix(in.Data, "\n页数：") {
						pagecount := strings.TrimPrefix(in.Data, "\n页数：")
						pagecount = strings.TrimRight(pagecount, "\n")

						var err error
						meta.PageCount, err = strconv.Atoi(pagecount)
						if err != nil {
							fmt.Println("Pagecount convert:", err)
						}
					}
					// fmt.Printf("in node:\n%+v\n", in)
					for ks := in.FirstChild; ks != nil; ks = ks.NextSibling {

						if strings.HasPrefix(ks.Data, "上架日期 : ") {
							content := strings.TrimPrefix(ks.Data, "上架日期 : ")
							meta.JMUploadTime = formatTime(content)

						}
						if strings.HasPrefix(ks.Data, "更新日期 : ") {
							content := strings.TrimPrefix(ks.Data, "更新日期 : ")
							meta.JMUpdateTime = formatTime(content)
						}

						if ok, id := nodeHasAttr(ks, "id"); ok {
							fmt.Println("id val:", id)
							for as := ks.FirstChild; as != nil; as = as.NextSibling {
								// fmt.Printf("as node\n%+v\n", as)
								if as.Data != "" {
									// fmt.Println("Album liked:", as.Data)
									count := strings.ToUpper(as.Data)
									if strings.HasSuffix(count, "K") {
										count = count[:len(count)-1] + "000"
									}
									meta.JMLiked, _ = strconv.Atoi(count)
								}

							}
						}

						for as := ks.FirstChild; as != nil; as = as.NextSibling {
							if as.FirstChild != nil {
								// fmt.Println("Album Viewd:", as.FirstChild.Data)
								count := strings.ToUpper(as.FirstChild.Data)
								if strings.HasSuffix(count, "K") {
									count = count[:len(count)-1] + "000"
								}
								meta.JMViewCount, _ = strconv.Atoi(count)
							}
						}

					}

				}
			}
		} else if nodeContainClass(node, "tag-block") {
			for in := node.FirstChild; in != nil; in = in.NextSibling {
				if ok, dataType := nodeHasAttr(in, "data-type"); ok {
					var tags []tagInfo
					switch dataType {
					case "works":
						// 作品
						fmt.Println("作品")
						tags = tagInfoInnerGet(in)
					case "actor":
						// 登场人物
						fmt.Println("登场人物")
						tags = tagInfoInnerGet(in)
						for _, v := range tags {
							link, _ := url.JoinPath(OrangeAppCCURLBase, v.Link)
							meta.Actor = append(meta.Actor, dao.ComicActor{
								Name:       v.Name,
								SearchName: v.Name,
								Link:       link,
							})
						}

					case "tags":
						// 标签
						fmt.Println("标签", dataType)
						tags = tagInfoInnerGet(in)
						for _, v := range tags {
							link, _ := url.JoinPath(OrangeAppCCURLBase, v.Link)
							meta.Tag = append(meta.Tag, dao.ComicTag{
								Name:       v.Name,
								SearchName: v.Name,
								Link:       link,
							})
						}
					case "author":
						// 作者
						fmt.Println("作者", dataType)
						tags = tagInfoInnerGet(in)
						for _, v := range tags {
							link, _ := url.JoinPath(OrangeAppCCURLBase, v.Link)
							meta.Artist = append(meta.Artist, dao.ComicAuthor{
								Name:       v.Name,
								SearchName: v.Name,
								Link:       link,
							})
						}
					}

				}
			}
		}

	}
}

type PosterInfo struct {
	PosterUrl string
}

func comicPosterFind(panel *goquery.Selection) (info PosterInfo) {

	albumCover := panel.Find(".col-lg-5")
	if len(albumCover.Nodes) == 0 {
		return PosterInfo{}
	}
	fmt.Println("albumCover nodes:", len(albumCover.Nodes))
	for _, v := range albumCover.Nodes {
		fmt.Printf("col-lg-5 nodes: %+v\n", v)
	}

	overlay := albumCover.Find(".thumb-overlay")
	fmt.Println("overlay nodes:", len(overlay.Nodes))
	for _, v := range overlay.Nodes {
		fmt.Printf("overlaynodes: %+v\n", v)
	}

	overlayImg := overlay.Find("img")
	fmt.Println("overlayImg nodes:", len(overlayImg.Nodes))
	for _, v := range overlayImg.Nodes {
		fmt.Printf("overlayImg nodes: %+v\n", v)
	}
	for _, img := range overlayImg.Nodes {
		for _, at := range img.Attr {
			if at.Key == "src" {
				fmt.Println("Overlay image:", at.Val)
				info.PosterUrl = at.Val
			}
		}
		// for node := img.FirstChild; node != nil; node = node.NextSibling {
		// 	for _, node := range overlayImg.Nodes {
		// 		fmt.Printf("%+v\n", node)
		// 		for _, at := range node.Attr {
		// 			if at.Key == "src" {
		// 				fmt.Println("Overlay image:", at.Val)
		// 				info.PosterUrl = at.Val
		// 			}
		// 		}
		// 	}
		// }
	}

	return
}

func rodGetHtml(requestURL string) (string, error) {
	url, err := launcher.New().Set("proxy-server", webbody.ProxyUrl).
		Delete("use-mock-keychain").
		Launch()
	if err != nil {
		return "", err
	}

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.Close()

	page := browser.MustPage(requestURL)
	s, err := page.HTML()
	if err != nil {
		return "", err
	}
	fmt.Println(s)
	return s, nil
}

func webbodyGetHtml(requestURL string) (string, error) {
	err, body := webbody.WebGet("GET", requestURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(string(body))
	return string(body), nil
}

func collyGetHtml(requestURL string) (html string, err error) {

	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher(webbody.ProxyUrl)
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Proxy Address: %s\n", r.Request.ProxyURL)
		// log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1)

		// 获取响应的 HTML
		html = string(r.Body)
		fmt.Println(html)
	})

	// Fetch httpbin.org/ip five times
	err = c.Visit(requestURL)
	if err != nil {
		return "", err
	}

	return
}

func gezGetHtml(requestURL string) (html string, err error) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{requestURL},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(string(r.Body))
			html = string(r.Body)
		},
		ProxyFunc: client.RoundRobinProxy(webbody.ProxyUrl),
		// BrowserEndpoint: "ws://localhost:3000",
	}).Start()
	return
}

func chromeDpGetHtml(requestURL string) (html string, err error) {
	chromeOpts := []chromedp.ExecAllocatorOption{
		chromedp.ProxyServer(webbody.ProxyUrl),
		// chromedp.Flag("headless", true),
		// chromedp.Flag("disable-gpu", true),
		// chromedp.Flag("mute-audio", true),
	}
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), chromeOpts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// // create chrome instance
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	chromedp.WithLogf(log.Printf),
	// )

	// defer cancel()

	// create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()
	start := time.Now()
	// navigate to a page, wait for an element, click
	// var res string

	chromedp.ProxyServer(webbody.ProxyUrl)
	err = chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(requestURL),
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.ScrollIntoView(`footer`),
		// chromedp.WaitVisible(`footer < div`),
		// chromedp.Text(`h1`, &res, chromedp.NodeVisible, chromedp.ByQuery),

		// chromedp.Text("html", &html, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.WaitReady(`.col-md-12`),
		chromedp.OuterHTML("html", &html),

		// chromedp.Text("html", &html, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if err != nil {
		return "", err
	}
	// fmt.Printf("h1 contains: '%s'\n", res)
	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())
	return
}

func OrangeAppDownload(requestURL string, savePath string, db *gorm.DB) error {
	// s, err := collyGetHtml(requestURL)
	// if err != nil {
	// 	return fmt.Errorf("Get html failed:%v\n", err)
	// }
	// fmt.Println(s)

	err, body := webbody.WebGet("GET", requestURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	meta, err := OrangeAppCrach(doc)
	if err != nil {
		fmt.Println(doc.Text())
		return err
	}
	meta.OriginURL = requestURL

	coverDir := filepath.Join(savePath, "cover")
	helpers.EnsurePath(coverDir)
	err, meta.PosterFilePath, _ = webbody.WriteFile("GET", meta.PosterURL, coverDir)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// fmt.Printf("%+v\n", meta)
	db.Save(&meta)

	return nil
}
