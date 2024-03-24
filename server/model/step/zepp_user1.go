// 自动生成模板ZeppUser1
package step

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zeppUser1表 结构体  ZeppUser1
type ZeppUser1 struct {
 global.GVA_MODEL 
      CreatorId  *int `json:"creatorId" form:"creatorId" gorm:"column:creator_id;comment:;size:19;"`  //创建者id 
      UserId  string `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:191;"`  //userId字段 
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;size:191;"`  //password字段 
      TargetStepCnt  *int `json:"targetStepCnt" form:"targetStepCnt" gorm:"column:target_step_cnt;comment:;size:19;"`  //targetStepCnt字段 
      TodayStepCnt  *int `json:"todayStepCnt" form:"todayStepCnt" gorm:"column:today_step_cnt;comment:;size:19;"`  //todayStepCnt字段 
      PushStatus  string `json:"pushStatus" form:"pushStatus" gorm:"column:push_status;comment:;size:191;"`  //pushStatus字段 
}


// TableName zeppUser1表 ZeppUser1自定义表名 zeppUser1
func (ZeppUser1) TableName() string {
  return "zeppUser1"
}

