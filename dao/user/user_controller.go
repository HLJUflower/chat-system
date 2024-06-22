package user

import "chat-system/config/global"

func (u User) CreateUser() error {
	err := global.Config.MysqlDB.Table("user").Create(&u).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}

func (u User) UpdateUser() error {
	err := global.Config.MysqlDB.Table("user").Save(&u).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}

func (u User) DeleteUser() error {
	err := global.Config.MysqlDB.Table("user").Delete(&u).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}

func (u User) GetUser() error {
	err := global.Config.MysqlDB.Table("user").First(&u).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}
