package taskmanager

import (
	"fmt"
	"sync"

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
	var countLocker sync.Mutex
	var onAir int
	var total int
	var worked int

	picChan := make(chan string, 3)
	go func() {
		for {
			select {
			case v := <-picChan:
				pic.Download(v, savePath)
				// time.Sleep(3 * time.Second)
				countLocker.Lock()
				onAir--
				worked++
				fmt.Printf("--> onAir:%d\tfinished:%d\ttotal:%d\n", onAir, worked, total)
				countLocker.Unlock()
			}
		}
	}()
	comicChan := make(chan string, 3)
	go func() {
		for {
			select {
			case v := <-comicChan:
				comic.Download(v, savePath)
				countLocker.Lock()
				onAir--
				worked++
				fmt.Printf("--> onAir:%d\tfinished:%d\ttotal:%d\n", onAir, worked, total)
				countLocker.Unlock()
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
				var legal bool
				if pic.IsCorrectUrl(v) {
					go func() {
						picChan <- v
					}()
					legal = true
				} else if comic.IsCorrectUrl(v) {
					go func() {
						comicChan <- v
					}()
					legal = true
				} else {
					fmt.Println("Task garbage:", v)
				}

				if legal {
					countLocker.Lock()
					onAir++
					total++
					fmt.Printf("--> onAir:%d\n", onAir)
					countLocker.Unlock()
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
