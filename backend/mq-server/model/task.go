package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null;column:uid"`
	Title     string `gorm:"index;not null;column:title"`
	Status    int    `gorm:"default:0;column:status"`
	Content   string `gorm:"type:longtext;column:content"`
	StartTime int64  `gorm:"default:'0'; column:start_time"`
	EndTime   int64  `gorm:"default:'0'; column:end_time"`
}
