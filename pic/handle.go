package pic

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"durian/dao"
)

type TagInfo struct {
	Name       string
	SearchName string
	Href       string
}

func Taginfo2DBCopyRight(src []TagInfo) (dst []dao.PictureCopyRight) {
	for _, v := range src {
		dst = append(dst, dao.PictureCopyRight{
			CopyRightName: v.Name,
			SearchName:    v.SearchName,
			Link:          v.Href,
		})
	}
	return
}

func Taginfo2DBArtist(src []TagInfo) (dst []dao.PictureArtist) {
	for _, v := range src {
		dst = append(dst, dao.PictureArtist{
			ArtistName: v.Name,
			SearchName: v.SearchName,
			Link:       v.Href,
		})
	}
	return
}

func Taginfo2DBMetaData(src []TagInfo) (dst []dao.PictureMetaData) {
	for _, v := range src {
		dst = append(dst, dao.PictureMetaData{
			MetaValue:  v.Name,
			SearchName: v.SearchName,
			Link:       v.Href,
		})
	}
	return
}

func Taginfo2DBTag(src []TagInfo) (dst []dao.PictureTag) {
	for _, v := range src {
		dst = append(dst, dao.PictureTag{
			TagName:    v.Name,
			SearchName: v.SearchName,
			Link:       v.Href,
		})
	}
	return
}

func Taginfo2DBCharacter(src []TagInfo) (dst []dao.PictureCharacter) {
	for _, v := range src {
		dst = append(dst, dao.PictureCharacter{
			CharacterName: v.Name,
			SearchName:    v.SearchName,
			Link:          v.Href,
		})
	}
	return
}

type TagMulti struct {
	Artist    []TagInfo
	MetaData  []TagInfo
	CopyRight []TagInfo
	Character []TagInfo
	General   []TagInfo
}

func IsCorrectUrl(requestUrl string) bool {
	ur, _ := url.ParseRequestURI(requestUrl)
	var ret bool
	switch ur.Host {
	case "gelbooru.com":
		ret = true
	case "yande.re":
		ret = true
	case "rule34.xxx":
		ret = true
	}
	return ret
}

func Download(requestUrl string, savePath string) {
	// fmt.Println("Download pic:", requestUrl)
	spi := strings.Split(requestUrl, "://")
	requestBody := spi[1]
	rq := strings.Split(requestBody, "/")
	var err error
	switch rq[0] {
	case "gelbooru.com":
		if strings.Contains(requestBody, "id=") {
			err = GelbooruDownloadPics(requestUrl, savePath, dao.Object)
		} else {
			err = fmt.Errorf("Not a gelbooru image page:%s\n", requestUrl)
		}

	case "yande.re":
		if strings.HasPrefix(requestBody, "yande.re/post/show") {
			err = YandeDownloadPics(requestUrl, savePath, dao.Object)
		} else {
			err = fmt.Errorf("Not a yande.re image page:%s\n", requestUrl)
		}
	case "rule34.xxx":
		if strings.Contains(requestBody, "id=") {
			err = Rule34DownloadPics(requestUrl, savePath, dao.Object)
		} else {
			err = fmt.Errorf("Not a rule34 image page:%s\n", requestUrl)
		}
	default:
		err = errors.New("no branch caughted:" + rq[0])
	}

	if err != nil {
		fmt.Println("Encounter error at downloading pics:", err)
	} else {
		// fmt.Println("finshied ", requestUrl)
	}

}
