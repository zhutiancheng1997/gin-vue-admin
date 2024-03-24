package step

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/step"
    stepReq "github.com/flipped-aurora/gin-vue-admin/server/model/step/request"
)

type ZeppLifeUserService struct {
}

// CreateZeppLifeUser 创建zepplife用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppLifeUserService *ZeppLifeUserService) CreateZeppLifeUser(zeppLifeUser *step.ZeppLifeUser) (err error) {
	err = global.GVA_DB.Create(zeppLifeUser).Error
	return err
}

// DeleteZeppLifeUser 删除zepplife用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppLifeUserService *ZeppLifeUserService)DeleteZeppLifeUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&step.ZeppLifeUser{},"id = ?",ID).Error
	return err
}

// DeleteZeppLifeUserByIds 批量删除zepplife用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppLifeUserService *ZeppLifeUserService)DeleteZeppLifeUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]step.ZeppLifeUser{},"id in ?",IDs).Error
	return err
}

// UpdateZeppLifeUser 更新zepplife用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppLifeUserService *ZeppLifeUserService)UpdateZeppLifeUser(zeppLifeUser step.ZeppLifeUser) (err error) {
	err = global.GVA_DB.Save(&zeppLifeUser).Error
	return err
}

// GetZeppLifeUser 根据ID获取zepplife用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppLifeUserService *ZeppLifeUserService)GetZeppLifeUser(ID string) (zeppLifeUser step.ZeppLifeUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&zeppLifeUser).Error
	return
}

// GetZeppLifeUserInfoList 分页获取zepplife用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (zeppLifeUserService *ZeppLifeUserService)GetZeppLifeUserInfoList(info stepReq.ZeppLifeUserSearch) (list []step.ZeppLifeUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&step.ZeppLifeUser{})
    var zeppLifeUsers []step.ZeppLifeUser
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
	
	err = db.Find(&zeppLifeUsers).Error
	return  zeppLifeUsers, total, err
}
