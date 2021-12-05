package model

import (
	"gorm.io/gorm"
)

type PedesFlowInfo struct {
	gorm.Model
	CameraID  uint   `gorm:"not null"`                  // 相机ID
	Time      string `gorm:"type:varchar(31);not null"` // 时间，格式为YYYYmmddHHMMSS，e.g. 194910010000
	PersonNum uint   `gorm:"not null"`                  // 人数
}
