// 自动生成模板ZeppUser
package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zeppUser 结构体  ZeppUser
type ZeppUser struct {
 global.GVA_MODEL 
      User  string `json:"user" form:"user" gorm:"column:user;comment:;"`  //user 
}


// TableName zeppUser ZeppUser自定义表名 zeppUser
func (ZeppUser) TableName() string {
  return "zeppUser"
}

