package taskmanager

import (
	"fmt"

	"durian/comic"
	"durian/pic"
	"github.com/jinzhu/gorm"
)

type TaskManager struct {
	Db        *gorm.DB
	SavePath  string
	TaskChan  chan string
	picChan   chan string
	comicChan chan string
}

func NewTaskManager(db *gorm.DB, savePath string) *TaskManager {
	picChan := make(chan string, 3)
	go func() {
		for {
			select {
			case v := <-picChan:
				pic.Download(v, savePath)
				// time.Sleep(3 * time.Second)
			}
		}
	}()
	comicChan := make(chan string, 3)
	go func() {
		for {
			select {
			case v := <-comicChan:
				comic.Download(v, savePath)
				// time.Sleep(3 * time.Second)
			}
		}
	}()
	TaskChan := make(chan string, 10)
	go func() {
		for {
			select {
			case v := <-TaskChan:
				fmt.Println("Task recved:", v)
				if pic.IsCorrectUrl(v) {
					go func() {
						picChan <- v
					}()
				} else if comic.IsCorrectUrl(v) {
					go func() {
						comicChan <- v
					}()
				} else {
					fmt.Println("Task garbage:", v)
				}

			}
		}

	}()

	return &TaskManager{
		Db:        db,
		SavePath:  savePath,
		TaskChan:  TaskChan,
		picChan:   picChan,
		comicChan: comicChan,
	}

}

func (t *TaskManager) AddTask(requestUrl string) {
	t.TaskChan <- requestUrl
}
