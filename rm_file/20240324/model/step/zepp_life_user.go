// 自动生成模板ZeppLifeUser
package step

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zepplife用户 结构体  ZeppLifeUser
type ZeppLifeUser struct {
 global.GVA_MODEL 
      User  string `json:"user" form:"user" gorm:"column:user;comment:;"`  //用户id 
}


// TableName zepplife用户 ZeppLifeUser自定义表名 zepp_life_user
func (ZeppLifeUser) TableName() string {
  return "zepp_life_user"
}

