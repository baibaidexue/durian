package pic

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"durian/dao"
	"durian/webbody"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/html"
)

const (
	GelbooruURLBase = "https://gelbooru.com/"
)

func GelbooruDownloadPics(requestURL string, savePath string, db *gorm.DB) error {
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

	// 检测是否是展示页面
	if sec := doc.Find("container"); sec != nil {

	}

	tag := GelbooruTagFind(doc)
	// fmt.Printf("%+v\n", tag)

	imageURL := GelbooruImageUrlFind(doc)
	imageURL = GelbooruImageUrlOrigin(imageURL)
	// fmt.Println("ImageURL:", imageURL)
	if imageURL == "" {
		return errors.New("Imageurl not found.")
	}

	err, s, i := webbody.WriteFile("GET", imageURL, savePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// fmt.Println("s:", s)
	// fmt.Println("bytes:", i)

	this := dao.Picture{
		OriginURL:  requestURL,
		Name:       "",
		SamplePath: "",
		ImageURL:   imageURL,
		FilePath:   s,
		FileName:   filepath.Base(s),
		Size:       uint64(i),
		Artist:     Taginfo2DBArtist(tag.Artist),
		Tags:       Taginfo2DBTag(tag.General),
		Character:  Taginfo2DBCharacter(tag.Character),
		Metadata:   Taginfo2DBMetaData(tag.MetaData),
		CopyRight:  Taginfo2DBCopyRight(tag.CopyRight),
	}

	// fmt.Printf("%+v\n", this)
	db.Save(&this)

	return nil
}

func GelbooruImageUrlOrigin(current string) string {
	temp := strings.ReplaceAll(current, "sample_", "")
	return strings.ReplaceAll(temp, "samples", "images")
}

func GelbooruImageUrlFind(doc *goquery.Document) (imageURL string) {
	doc.Find(".aside").Each(func(i int, s *goquery.Selection) {
		// fmt.Printf("Aside: %+v\n", s)
		li := s.Find("a")
		for _, v := range li.Nodes {

			// fmt.Println("_______________")
			// fmt.Printf("Aside a nodes: %+v\n", d)
			valNo := v.FirstChild
			if valNo != nil {
				if valNo.Data == "Original image" {
					for _, at := range v.Attr {
						if at.Key == "href" {
							imageURL = at.Val
							return
						}
						// fmt.Println(at.Key, ":", at.Val)
					}
				}
			}
			// fmt.Println("_______________")
		}
	})

	if imageURL != "" {
		return
	}

	selection := doc.Find("picture")
	for _, node := range selection.Nodes {
		// fmt.Printf("node: %+v\n", node)
		var depth int
		var imgNodes []*html.Node
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			func(n *html.Node, depth int) {
				// fmt.Printf("child n: %+v\n", n)
				if n.Data == "img" {
					imgNodes = append(imgNodes, n)
				}
			}(child, depth+1)
		}

		for _, n := range imgNodes {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					imageURL = attr.Val
					return
				}
			}
		}
		// fmt.Println("lenImageNodes:", len(imgNodes))
	}

	return
}

func GelbooruTagFind(doc *goquery.Document) (geltag TagMulti) {
	doc.Find(".tag-list").Each(func(i int, selection *goquery.Selection) {
		for _, node := range selection.Nodes {
			var depth int

			var tagNodes []struct {
				Class string
				Node  *html.Node
			}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				func(n *html.Node, depth int) {
					var nodeType string
					for _, attr := range n.Attr {
						if attr.Key == "class" {
							nodeType = attr.Val
						}
					}
					switch nodeType {
					case "tag-type-artist":
						fallthrough
					case "tag-type-character":
						fallthrough
					case "tag-type-metadata":
						fallthrough
					case "tag-type-general":
						tagNodes = append(tagNodes, struct {
							Class string
							Node  *html.Node
						}{Class: nodeType, Node: n})
					case "sm-hidden":
					}
				}(child, depth+1)
			}

			for _, v := range tagNodes {
				var val string
				var link string
				for child := v.Node.FirstChild; child != nil; child = child.NextSibling {
					if child.Data == "a" && child.Type == html.ElementNode {
						valnode := child.FirstChild
						if valnode == nil {
							fmt.Println("a had a nil node")
						} else {
							val = valnode.Data
							for _, at := range child.Attr {
								if at.Key == "href" {
									link = at.Val
								}
							}

						}
					}
				}

				if val == "" || link == "" {
					continue
				}
				// link = GelbooruURLBase + link
				link, _ = url.JoinPath(GelbooruURLBase, link)
				switch v.Class {
				case "tag-type-artist":
					si := strings.Split(val, "=")
					geltag.Artist = append(geltag.Artist, TagInfo{
						Name:       val,
						SearchName: si[len(si)-1],
						Href:       link,
					})

				case "tag-type-character":
					si := strings.Split(val, "=")
					geltag.Character = append(geltag.Character, TagInfo{
						Name:       val,
						SearchName: si[len(si)-1],
						Href:       link,
					})
				case "tag-type-copyright":
					si := strings.Split(val, "=")
					geltag.CopyRight = append(geltag.CopyRight, TagInfo{
						Name:       val,
						SearchName: si[len(si)-1],
						Href:       link,
					})

				case "tag-type-metadata":
					si := strings.Split(val, "=")
					geltag.MetaData = append(geltag.MetaData, TagInfo{
						Name:       val,
						SearchName: si[len(si)-1],
						Href:       link,
					})

				case "tag-type-general":
					si := strings.Split(val, "=")
					geltag.General = append(geltag.General, TagInfo{
						Name:       val,
						SearchName: si[len(si)-1],
						Href:       link,
					})

				}
			}

		}

	})
	return
}
