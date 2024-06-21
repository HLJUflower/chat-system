package sql

import (
	"chat-system/utils/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:zl20040203@(localhost:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		logging.Error("DB connect failed : ", err)
		return nil
	}
	return db
}
