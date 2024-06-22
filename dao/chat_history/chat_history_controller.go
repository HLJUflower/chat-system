package chat_history

import "chat-system/config/global"

//func (h ChatHistory) CreateChatHistory() error {
//	connect := global.Config.MongoDB.Database("chat").Collection("chat_history")
//	_, err := connect.InsertOne(context.Background(), h)
//	if err != nil {
//		global.Config.Logger.Error(err)
//	}
//	return err
//}

func (h ChatHistory) CreateChatHistory() error {
	err := global.Config.MysqlDB.Table("chat_history").Create(&h).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}

func (h ChatHistory) UpdateChatHistory() error {
	err := global.Config.MysqlDB.Table("chat_history").Save(&h).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}

func (h ChatHistory) GetChatHistory() error {
	err := global.Config.MysqlDB.Table("chat_history").First(&h).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}

func (h ChatHistory) DeleteChatHistory() error {
	err := global.Config.MysqlDB.Table("chat_history").Delete(&h).Error
	if err != nil {
		global.Config.Logger.Error(err)
	}
	return err
}
