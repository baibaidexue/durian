package dao

import "github.com/jinzhu/gorm"

type AirTask struct {
	gorm.Model
	OriginURL string `gorm:"unique_index:unique_idx_air_task_url;column:origin_url"`
	ErrorInfo string `gorm:"column:error_info"`
}
