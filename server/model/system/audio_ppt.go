
// 自动生成模板AudioPpt
package system
import (
	"time"
)

// audioPpt表 结构体  AudioPpt
type AudioPpt struct {
  Id  *int32 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //id字段
  Uid  *int32 `json:"uid" form:"uid" gorm:"comment:用户id;column:uid;"`  //用户id
  TaskId  *string `json:"taskId" form:"taskId" gorm:"comment:任务唯一识别ID;column:task_id;size:50;"`  //任务唯一识别ID
  AudioId  *int32 `json:"audioId" form:"audioId" gorm:"comment:用户id;column:audio_id;"`  //用户id
  SummaryId  *string `json:"summaryId" form:"summaryId" gorm:"comment:总结结果data_id;column:summary_id;size:50;"`  //总结结果data_id
  Style  *bool `json:"style" form:"style" gorm:"comment:PPT 风格;column:style;"`  //PPT 风格
  DetailLevel  *string `json:"detailLevel" form:"detailLevel" gorm:"comment:内容详略级别;column:detail_level;size:40;"`  //内容详略级别
  SlideNum  *string `json:"slideNum" form:"slideNum" gorm:"comment:页数档位;column:slide_num;size:40;"`  //页数档位
  Outline  *string `json:"outline" form:"outline" gorm:"comment:分页提纲 JSON；如果 slide_num 变化，则 outline 需要重新生成;column:outline;type:text;"`  //分页提纲 JSON；如果 slide_num 变化，则 outline 需要重新生成
  Provider  *string `json:"provider" form:"provider" gorm:"comment:模型供应商;column:provider;size:100;"`  //模型供应商
  Status  *int32 `json:"status" form:"status" gorm:"comment:生成状态,300:生成中;column:status;"`  //生成状态,300:生成中
  Imgs  *string `json:"imgs" form:"imgs" gorm:"comment:图片结果,保存对象存储的path,多张图片逗号分隔;column:imgs;size:1000;"`  //图片结果,保存对象存储的path,多张图片逗号分隔
  Ppt  *string `json:"ppt" form:"ppt" gorm:"comment:幻灯片结果,保存对象存储的path,仅一个;column:ppt;size:100;"`  //幻灯片结果,保存对象存储的path,仅一个
  Pdf  *string `json:"pdf" form:"pdf" gorm:"comment:pdf结果,保存对象存储的path,仅一个;column:pdf;size:100;"`  //pdf结果,保存对象存储的path,仅一个
  SlideCount  *int32 `json:"slideCount" form:"slideCount" gorm:"comment:幻灯片数量;column:slide_count;"`  //幻灯片数量
  AnalysisTokens  *int32 `json:"analysisTokens" form:"analysisTokens" gorm:"comment:analysis tokens数;column:analysis_tokens;"`  //analysis tokens数
  ImageTokens  *int32 `json:"imageTokens" form:"imageTokens" gorm:"comment:image tokens数;column:image_tokens;"`  //image tokens数
  TotalTokens  *int32 `json:"totalTokens" form:"totalTokens" gorm:"comment:总tokens数;column:total_tokens;"`  //总tokens数
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:创建时间;column:created_at;"`  //创建时间
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间;column:updated_at;"`  //更新时间
  DeletedAt  *int32 `json:"deletedAt" form:"deletedAt" gorm:"comment:删除写入时间戳;column:deleted_at;"`  //删除写入时间戳
}


// TableName audioPpt表 AudioPpt自定义表名 audio_ppt
func (AudioPpt) TableName() string {
    return "audio_ppt"
}





