// 自动生成模板Test1User
package zeppLife

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Test1User 结构体  Test1User
type Test1User struct {
 global.GVA_MODEL 
      User  string `json:"user" form:"user" gorm:"column:user;comment:;"`  //user 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;"`  //password 
      PushType  *int `json:"pushType" form:"pushType" gorm:"column:push_type;comment:;"`  //刷步方式 
      TodayStepCnt  *int `json:"todayStepCnt" form:"todayStepCnt" gorm:"column:today_step_cnt;comment:;"`  //今日步数 
}


// TableName Test1User Test1User自定义表名 Test1User
func (Test1User) TableName() string {
  return "Test1User"
}

