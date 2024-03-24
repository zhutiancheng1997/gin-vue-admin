package step

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/step"
    stepReq "github.com/flipped-aurora/gin-vue-admin/server/model/step/request"
)

type ZeppUser1Service struct {
}

// CreateZeppUser1 创建zeppUser1表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUser1Service *ZeppUser1Service) CreateZeppUser1(zeppUser1 *step.ZeppUser1) (err error) {
	err = global.GVA_DB.Create(zeppUser1).Error
	return err
}

// DeleteZeppUser1 删除zeppUser1表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUser1Service *ZeppUser1Service)DeleteZeppUser1(ID string) (err error) {
	err = global.GVA_DB.Delete(&step.ZeppUser1{},"id = ?",ID).Error
	return err
}

// DeleteZeppUser1ByIds 批量删除zeppUser1表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUser1Service *ZeppUser1Service)DeleteZeppUser1ByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]step.ZeppUser1{},"id in ?",IDs).Error
	return err
}

// UpdateZeppUser1 更新zeppUser1表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUser1Service *ZeppUser1Service)UpdateZeppUser1(zeppUser1 step.ZeppUser1) (err error) {
	err = global.GVA_DB.Save(&zeppUser1).Error
	return err
}

// GetZeppUser1 根据ID获取zeppUser1表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUser1Service *ZeppUser1Service)GetZeppUser1(ID string) (zeppUser1 step.ZeppUser1, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&zeppUser1).Error
	return
}

// GetZeppUser1InfoList 分页获取zeppUser1表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUser1Service *ZeppUser1Service)GetZeppUser1InfoList(info stepReq.ZeppUser1Search) (list []step.ZeppUser1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&step.ZeppUser1{})
    var zeppUser1s []step.ZeppUser1
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&zeppUser1s).Error
	return  zeppUser1s, total, err
}
