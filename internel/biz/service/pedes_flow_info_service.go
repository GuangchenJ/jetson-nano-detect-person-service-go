package service

import (
	"gorm.io/gorm"

	"go_server/internel/data/model"
)

// PedesFlowService 人流量服务
type PedesFlowService interface {
	// NewPedesFlowInfo 添加新的图书馆区域信息
	NewPedesFlowInfo(pedesFlowInfo model.PedesFlowInfo) error
}

// 新增信息服务的实现结构体
type pedesFlowService struct {
	db *gorm.DB
}

func NewPedesFlowService(db *gorm.DB) PedesFlowService {
	return &pedesFlowService{
		db: db,
	}
}

func (pc *pedesFlowService) NewPedesFlowInfo(pedesFlowInfo model.PedesFlowInfo) error {
	// 	pedesFlowInfo.Time = time.Now().Format("20060102150405")
	result := pc.db.Create(&pedesFlowInfo)
	if nil != result.Error {
		return result.Error
	}
	//     fmt.Println(result.RowsAffected) // 返回插入记录的条数
	return nil
}
