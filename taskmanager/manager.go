package taskmanager

import (
	"fmt"
	"sync"
	"time"

	"durian/comic"
	"durian/dao"
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
	failed      int
	worked      int
}

func comicRun(tm *TaskManager) {
	for {
		select {
		case taskUrl := <-tm.comicChan:
			err := comic.Download(taskUrl, tm.SavePath)
			tm.endEvent(taskUrl, err)
		}
	}
}

func picRun(tm *TaskManager) {
	for {
		select {
		case taskUrl := <-tm.picChan:
			err := pic.Download(taskUrl, tm.SavePath)
			tm.endEvent(taskUrl, err)
		}
	}
}

func (tm *TaskManager) run() {

	go picRun(tm)
	go comicRun(tm)

	var alarmTime time.Duration
	alarmTime = 3 * time.Minute
	tick := time.NewTicker(alarmTime)
	var lastAirTaskId uint
	for {
		select {
		case <-tick.C:
			// load 5 failed task
			var airTaskList []dao.AirTask
			tm.Db.Where("id > ?", lastAirTaskId).Order("id ASC").Limit(5).Scan(&airTaskList)
			for _, v := range airTaskList {
				if v.ID > lastAirTaskId {
					lastAirTaskId = v.ID
				}

				tm.AddTask(v.OriginURL)
			}

			tick.Reset(alarmTime)
		case taskUrl := <-tm.TaskChan:

			lech := taskUrlValidate(taskUrl)
			if lech == "" {
				continue
			}

			fmt.Println("Task recved:", taskUrl)
			tm.startEvent(taskUrl)
			switch lech {
			case "pic":
				go func() {
					tm.picChan <- taskUrl
				}()
			case "comic":
				go func() {
					tm.comicChan <- taskUrl
				}()
			default:
			}

		}
	}
}

func taskUrlValidate(taskUrl string) (lech string) {
	if pic.IsCorrectUrl(taskUrl) {
		return "pic"
	}
	if comic.IsCorrectUrl(taskUrl) {
		return "comic"
	}
	return
}

func NewTaskManager(db *gorm.DB, savePath string) *TaskManager {
	tm := &TaskManager{
		Db:        db,
		SavePath:  savePath,
		TaskChan:  make(chan string, 60),
		picChan:   make(chan string, 5),
		comicChan: make(chan string, 5),
	}

	go tm.run()

	return tm
}

func (t *TaskManager) AddTask(requestUrl string) {
	t.TaskChan <- requestUrl
}
