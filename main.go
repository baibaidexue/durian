package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"durian/dao"
	"durian/helpers"
	"durian/taskmanager"
	"durian/webbody"
	"github.com/atotto/clipboard"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ProxyURL    string `yaml:"proxyurl"`
	DownloadDir string `yaml:"downloaddir"`
	DBDir       string `yaml:"databasedir"`
}

func main() {
	configFile, err := ioutil.ReadFile("config.yaml")
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

	manager := taskmanager.NewTaskManager(dao.Object, config.DownloadDir)
	// manager.AddTask("https://yande.re/post/show/1094503")
	// manager.AddTask("https://gelbooru.com/index.php?page=post&s=view&id=8619847&tags=thighband_pantyhose")

	clipboardmsg(manager)
}

func clipboardmsg(manager *taskmanager.TaskManager) {
	// TODO 当前的 clipboard 非常容易崩溃
	prev, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("clipboard readall openeerr")
		panic(err)
	}

	for {
		// 每隔1秒钟检查一次剪贴板内容是否发生变化
		time.Sleep(1 * time.Second)
		curr, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("clipboard readall openeerr")
			panic(err)
		}

		// 如果发现剪贴板内容发生变化，则将其写入文件
		if curr != prev {
			if strings.HasPrefix(curr, "http") {
				manager.AddTask(curr)
			}
			prev = curr
		}
	}
}
