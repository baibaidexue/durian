package comic

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

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

func nodeHasAttr(node *html.Node, attr string) bool {
	for _, v := range node.Attr {
		if v.Key == attr {
			return true
		}
	}
	return false
}

func whatMetaNode(node *html.Node) string {
	var count int
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		count++
	}

	// total 10 with tag block
	if count >= 8 {
		return "tag-block"
	}

	if nodeContainClass(node, "p-t-5 p-b-5 read-block") {
		return "read-block"
	}
	return "unhandled"
}

func OrangeAppMetaFind(doc *goquery.Document) {
	panel := doc.Find(".col-lg-7")
	fmt.Println("panel nodes:", len(panel.Nodes))

	for _, b := range panel.Nodes {
		for node := b.FirstChild; node != nil; node = node.NextSibling {

			switch whatMetaNode(node) {
			case "tag-block":
				fmt.Println("tag block:")
				if nodeContainClass(node, "") {
					fmt.Printf(" %+v\n", node)
				}
				dealTagInfo(node)

			case "read-block":
			}

		}
	}

}

func dealTagInfo(b *html.Node) {
	for node := b.FirstChild; node != nil; node = node.NextSibling {
		if nodeContainClass(node, "p-t-5 p-b-5") {
			if node.Type == html.TextNode {
				fmt.Printf("text node: %+v\n", node)

			}
			if node.Type == html.ElementNode {
				fmt.Printf("element node: %+v\n", node)
			}
		} else if nodeContainClass(node, "tag-block") {
			for in := node.FirstChild; in != nil; in = in.NextSibling {
				fmt.Printf("tag-block node: %+v\n", node)
			}
		}

	}
}

type PosterInfo struct {
	PosterUrl string
}

func comicPosterFind(doc *goquery.Document) (info PosterInfo) {
	panel := doc.Find(".col-md-12")
	fmt.Println("panel nodes:", len(panel.Nodes))

	albumCover := panel.Find(".col-lg-5")
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
	s, err := chromeDpGetHtml(requestURL)
	if err != nil {
		return fmt.Errorf("Get html failed:%v\n", err)
	}
	fmt.Println(s)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
		return err
	}

	short := comicShortFind(doc)
	fmt.Printf("shortInfo: %+v\n", short)

	poster := comicPosterFind(doc)
	fmt.Printf("posterInfo: %+v\n", poster)

	// imageURL := GelbooruImageUrlFind(doc)
	// imageURL = GelbooruImageUrlOrigin(imageURL)
	// // fmt.Println("ImageURL:", imageURL)
	// if imageURL == "" {
	// 	return errors.New("Imageurl not found.")
	// }
	//
	// err, s, i := webbody.WriteFile("GET", imageURL, savePath)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// // fmt.Println("s:", s)
	// // fmt.Println("bytes:", i)
	//
	// this := dao.Picture{
	// 	OriginURL:  requestURL,
	// 	Name:       "",
	// 	SamplePath: "",
	// 	ImageURL:   imageURL,
	// 	FilePath:   s,
	// 	FileName:   filepath.Base(s),
	// 	Size:       uint64(i),
	// 	Artist:     Taginfo2DBArtist(tag.Artist),
	// 	Tags:       Taginfo2DBTag(tag.General),
	// 	Character:  Taginfo2DBCharacter(tag.Character),
	// 	Metadata:   Taginfo2DBMetaData(tag.MetaData),
	// 	CopyRight:  Taginfo2DBCopyRight(tag.CopyRight),
	// }
	//
	// // fmt.Printf("%+v\n", this)
	// db.Save(&this)

	return nil
}
