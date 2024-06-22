package user

import "chat-system/config/global"

func (u User) CreateUser(user User) string {
	err := global.Config.DB.Table("users").Create(&user).Error
	if err != nil {
		global.Config.Logger.Error(err)
		return err.Error()
	}
	return "ok"
}

func (u User) UpdateUser(user User) string {
	err := global.Config.DB.Table("users").Save(&user).Error
	if err != nil {
		global.Config.Logger.Error(err)
		return err.Error()
	}
	return "ok"
}

func (u User) DeleteUser(user User) string {
	err := global.Config.DB.Table("users").Delete(&user).Error
	if err != nil {
		global.Config.Logger.Error(err)
		return err.Error()
	}
	return "ok"
}

func (u User) GetUser(user User) string {
	err := global.Config.DB.Table("users").First(&user).Error
	if err != nil {
		global.Config.Logger.Error(err)
		return err.Error()
	}
	return "ok"
}
