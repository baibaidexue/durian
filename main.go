package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"durian/dao"
	"durian/helpers"
	"durian/taskmanager"
	"durian/webbody"
	clib "golang.design/x/clipboard"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ProxyURL    string `yaml:"proxyurl"`
	DownloadDir string `yaml:"downloaddir"`
	DBDir       string `yaml:"databasedir"`
	Comicdir    string `yaml:"comicdir"`
}

func main() {
	configFile, err := ioutil.ReadFile("config.yaml")
	// configFile, err := ioutil.ReadFile("test/config.yaml")
	if err != nil {
		panic(err)
	}

	// 解析YAML配置文件
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)

	webbody.WebSetProxy(config.ProxyURL)
	helpers.EnsurePath(config.DBDir)
	dao.NewDb(dao.DatabaseOptions{}, config.DBDir)

	helpers.EnsurePath(config.DownloadDir)

	manager := taskmanager.NewTaskManager(dao.Object, config.DownloadDir, config.Comicdir)
	// manager.AddTask("https://yande.re/post/show/1094503")
	// manager.AddTask("https://gelbooru.com/index.php?page=post&s=view&id=8619847&tags=thighband_pantyhose")
	clipboardListening(manager)
}

func clipboardListening(manager *taskmanager.TaskManager) {
	err := clib.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	ch := clib.Watch(context.TODO(), clib.FmtText)
	var curr, prev string
	fmt.Println("Watching on...")
	for data := range ch {
		curr = string(data)
		// 如果发现剪贴板内容发生变化，则将其写入文件
		if curr != prev {
			if strings.HasPrefix(curr, "http") {
				manager.AddTask(curr)
			}
			prev = curr
		}
	}
}
