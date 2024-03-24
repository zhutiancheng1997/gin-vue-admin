package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife"
    zeppLifeReq "github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife/request"
)

type XmUserService struct {
}

// CreateXmUser 创建xmUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (xmUserService *XmUserService) CreateXmUser(xmUser *zeppLife.XmUser) (err error) {
	err = global.GVA_DB.Create(xmUser).Error
	return err
}

// DeleteXmUser 删除xmUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (xmUserService *XmUserService)DeleteXmUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&zeppLife.XmUser{},"id = ?",ID).Error
	return err
}

// DeleteXmUserByIds 批量删除xmUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (xmUserService *XmUserService)DeleteXmUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]zeppLife.XmUser{},"id in ?",IDs).Error
	return err
}

// UpdateXmUser 更新xmUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (xmUserService *XmUserService)UpdateXmUser(xmUser zeppLife.XmUser) (err error) {
	err = global.GVA_DB.Save(&xmUser).Error
	return err
}

// GetXmUser 根据ID获取xmUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (xmUserService *XmUserService)GetXmUser(ID string) (xmUser zeppLife.XmUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&xmUser).Error
	return
}

// GetXmUserInfoList 分页获取xmUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (xmUserService *XmUserService)GetXmUserInfoList(info zeppLifeReq.XmUserSearch) (list []zeppLife.XmUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&zeppLife.XmUser{})
    var xmUsers []zeppLife.XmUser
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
	
	err = db.Find(&xmUsers).Error
	return  xmUsers, total, err
}
