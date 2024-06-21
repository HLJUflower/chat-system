package user

import "time"

type User struct {
	UserId            int64     `json:"user_id,omitempty" gorm:"user_id"`                         //用户ID
	UserName          string    `json:"user_name,omitempty" gorm:"user_name"`                     //用户名
	UserInformationId string    `json:"user_information_id,omitempty" gorm:"user_information_id"` //用户身份证号
	UserRealName      string    `json:"user_real_name,omitempty" gorm:"user_real_name"`           //用户真实姓名
	Phone             string    `json:"phone,omitempty" gorm:"phone"`                             //联系电话
	Status            bool      `json:"status,omitempty" gorm:"status"`                           //状态（1正常 0离线）
	DelFlag           bool      `json:"del_flag,omitempty" gorm:"del_flag"`                       //删除标志（1代表存在 0代表删除）
	CreateTime        time.Time `json:"create_time" gorm:"create_time"`                           //创建时间
	UpdateTime        time.Time `json:"update_time" gorm:"update_time"`                           //登录时间
	Remark            string    `json:"remark,omitempty" gorm:"remark"`                           //备注
}
