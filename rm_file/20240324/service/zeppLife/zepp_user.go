package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife"
    zeppLifeReq "github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife/request"
)

type ZeppUserService struct {
}

// CreateZeppUser 创建zeppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUserService *ZeppUserService) CreateZeppUser(zeppUser *zeppLife.ZeppUser) (err error) {
	err = global.GVA_DB.Create(zeppUser).Error
	return err
}

// DeleteZeppUser 删除zeppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUserService *ZeppUserService)DeleteZeppUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&zeppLife.ZeppUser{},"id = ?",ID).Error
	return err
}

// DeleteZeppUserByIds 批量删除zeppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUserService *ZeppUserService)DeleteZeppUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]zeppLife.ZeppUser{},"id in ?",IDs).Error
	return err
}

// UpdateZeppUser 更新zeppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUserService *ZeppUserService)UpdateZeppUser(zeppUser zeppLife.ZeppUser) (err error) {
	err = global.GVA_DB.Save(&zeppUser).Error
	return err
}

// GetZeppUser 根据ID获取zeppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUserService *ZeppUserService)GetZeppUser(ID string) (zeppUser zeppLife.ZeppUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&zeppUser).Error
	return
}

// GetZeppUserInfoList 分页获取zeppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppUserService *ZeppUserService)GetZeppUserInfoList(info zeppLifeReq.ZeppUserSearch) (list []zeppLife.ZeppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&zeppLife.ZeppUser{})
    var zeppUsers []zeppLife.ZeppUser
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
	
	err = db.Find(&zeppUsers).Error
	return  zeppUsers, total, err
}
