package chat_history

import (
	"time"
)

type ChatHistory struct {
	ChatID       string    `json:"chat_id" gorm:"chat_id"`
	RequestName  int64     `json:"request_name" gorm:"request_name"`     //发送者
	ResponseName int64     `json:"response_name" gorm:"response_name"`   //接收者
	Txt          string    `json:"txt,omitempty" gorm:"txt"`             //内容
	DelFlag      bool      `json:"del_flag,omitempty" gorm:"del_flag"`   //删除标志（1代表存在 0代表删除）
	CreateTime   time.Time `json:"create_time" gorm:"create_time"`       //创建时间
	CreateIp     string    `json:"create_ip,omitempty" gorm:"create_ip"` //发送IP
	Remark       string    `json:"remark,omitempty" gorm:"remark"`       //备注
}
