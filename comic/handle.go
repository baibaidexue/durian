package comic

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"durian/dao"
)

func IsCorrectUrl(requestUrl string) bool {
	// return false
	ur, _ := url.ParseRequestURI(requestUrl)
	var ret bool
	switch ur.Host {
	case "orangeapp.cc":
		ret = true
	case "18comic.org":
		ret = true
	}
	return ret
}

func Download(requestUrl string, savePath string) error {
	fmt.Println("Download comic meta:", requestUrl)
	spi := strings.Split(requestUrl, "://")
	requestBody := spi[1]
	rq := strings.Split(requestBody, "/")
	var err error
	switch rq[0] {
	case "orangeapp.cc":
		fallthrough
	case "18comic.vip":
		fallthrough
	case "18comic.org":
		if strings.Contains(requestBody, "/album/") {
			err = OrangeAppDownload(requestUrl, savePath, dao.Object)
		} else {
			err = fmt.Errorf("Not a 18comic image page:%s\n", requestUrl)
		}
	default:
		err = errors.New("no branch caughted:" + rq[0])
	}

	return err
}
