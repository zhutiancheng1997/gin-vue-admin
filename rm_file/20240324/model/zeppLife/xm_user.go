// 自动生成模板XmUser
package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// xmUser 结构体  XmUser
type XmUser struct {
 global.GVA_MODEL 
      User  string `json:"user" form:"user" gorm:"column:user;comment:;"`  //user 
}


// TableName xmUser XmUser自定义表名 xm_user
func (XmUser) TableName() string {
  return "xm_user"
}

