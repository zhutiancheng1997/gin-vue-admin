package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife"
    zeppLifeReq "github.com/flipped-aurora/gin-vue-admin/server/model/zeppLife/request"
)

type Test1UserService struct {
}

// CreateTest1User 创建Test1User记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestUserService *Test1UserService) CreateTest1User(TestUser *zeppLife.Test1User) (err error) {
	err = global.GVA_DB.Create(TestUser).Error
	return err
}

// DeleteTest1User 删除Test1User记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestUserService *Test1UserService)DeleteTest1User(ID string) (err error) {
	err = global.GVA_DB.Delete(&zeppLife.Test1User{},"id = ?",ID).Error
	return err
}

// DeleteTest1UserByIds 批量删除Test1User记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestUserService *Test1UserService)DeleteTest1UserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]zeppLife.Test1User{},"id in ?",IDs).Error
	return err
}

// UpdateTest1User 更新Test1User记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestUserService *Test1UserService)UpdateTest1User(TestUser zeppLife.Test1User) (err error) {
	err = global.GVA_DB.Save(&TestUser).Error
	return err
}

// GetTest1User 根据ID获取Test1User记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestUserService *Test1UserService)GetTest1User(ID string) (TestUser zeppLife.Test1User, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&TestUser).Error
	return
}

// GetTest1UserInfoList 分页获取Test1User记录
// Author [piexlmax](https://github.com/piexlmax)
func (TestUserService *Test1UserService)GetTest1UserInfoList(info zeppLifeReq.Test1UserSearch) (list []zeppLife.Test1User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&zeppLife.Test1User{})
    var TestUsers []zeppLife.Test1User
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
	
	err = db.Find(&TestUsers).Error
	return  TestUsers, total, err
}
