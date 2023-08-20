package taskmanager

import (
	"fmt"
	"os"

	"durian/dao"
)

func (t *TaskManager) startEvent(taskUrl string) {
	t.countLocker.Lock()
	defer t.countLocker.Unlock()

	t.onAir++
	t.total++
	fmt.Printf("--> onAir:%d\n", t.onAir)
	t.Db.Save(&dao.AirTask{
		OriginURL: taskUrl,
		ErrorInfo: "",
	})

}

func (t *TaskManager) endEvent(taskUrl string, err error) {
	t.countLocker.Lock()
	defer t.countLocker.Unlock()

	t.onAir--
	if err == nil {
		t.worked++
	} else {
		t.failed++
	}
	fmt.Printf("--> onAir:%d\tfailed:%d\tfinished:%d\ttotal:%d\n", t.onAir, t.failed, t.worked, t.total)
	if err == os.ErrExist {
		err = nil
	}

	if err != nil {
		t.Db.Model(&dao.AirTask{}).Where(map[string]interface{}{"origin_url": taskUrl}).Update(&dao.AirTask{
			ErrorInfo: err.Error(),
		})
	} else {
		t.Db.Model(&dao.AirTask{}).Unscoped().Where(map[string]interface{}{"origin_url": taskUrl}).Delete(&dao.AirTask{})
	}
}
