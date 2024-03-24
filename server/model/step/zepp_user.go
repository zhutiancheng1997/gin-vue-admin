// 自动生成模板ZeppUser
package step

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zeppUser 结构体  ZeppUser
type ZeppUser struct {
 global.GVA_MODEL 
      UserId  string `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`  //zepp用户id 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;"`  //密码 
      TargetStepCnt  *int `json:"target_step_cnt" form:"target_step_cnt" gorm:"column:target_step_cnt;comment:;"`  //目标步数 
      TodayStepCnt  *int `json:"todayStepCnt" form:"todayStepCnt" gorm:"column:today_step_cnt;comment:;"`  //今日步数 
      PushStatus  string `json:"pushStatus" form:"pushStatus" gorm:"column:push_status;comment:;"`  //刷步状态 
}


// TableName zeppUser ZeppUser自定义表名 zeppUser
func (ZeppUser) TableName() string {
  return "zeppUser"
}

