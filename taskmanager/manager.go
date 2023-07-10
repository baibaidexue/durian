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

	countLocker sync.Mutex
	onAir       int
	total       int
	worked      int
}

func (t *TaskManager) addCount() {
	t.countLocker.Lock()
	defer t.countLocker.Unlock()

	t.onAir++
	t.total++
	fmt.Printf("--> onAir:%d\n", t.onAir)

}

func (t *TaskManager) delCount() {
	t.countLocker.Lock()
	defer t.countLocker.Unlock()
	
	t.onAir--
	t.worked++
	fmt.Printf("--> onAir:%d\tfinished:%d\ttotal:%d\n", t.onAir, t.worked, t.total)

}

func comicRun(tm *TaskManager) {
	for {
		select {
		case v := <-tm.comicChan:
			comic.Download(v, tm.SavePath)
			// time.Sleep(3 * time.Second)
			tm.delCount()
		}
	}
}

func picRun(tm *TaskManager) {
	for {
		select {
		case v := <-tm.picChan:
			pic.Download(v, tm.SavePath)
			// time.Sleep(3 * time.Second)
			tm.delCount()
		}
	}
}

func span(tm *TaskManager) {
	for {
		select {
		case v := <-tm.TaskChan:
			fmt.Println("Task recved:", v)
			var legal bool
			if pic.IsCorrectUrl(v) {
				go func() {
					tm.picChan <- v
				}()
				legal = true
			} else if comic.IsCorrectUrl(v) {
				go func() {
					tm.comicChan <- v
				}()
				legal = true
			} else {
				fmt.Println("Task garbage:", v)
			}

			if legal {
				tm.addCount()
			}

		}
	}
}

func NewTaskManager(db *gorm.DB, savePath string) *TaskManager {
	tm := &TaskManager{
		Db:        db,
		SavePath:  savePath,
		TaskChan:  make(chan string, 10),
		picChan:   make(chan string, 3),
		comicChan: make(chan string, 3),
	}

	go picRun(tm)
	go comicRun(tm)

	go span(tm)

	return tm
}

func (t *TaskManager) AddTask(requestUrl string) {
	t.TaskChan <- requestUrl
}
