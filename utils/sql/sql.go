package sql

import (
	"chat-system/config/global"
	"chat-system/utils/logging"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:zl20040203@(localhost:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		logging.Error("MysqlDB connect failed : ", err)
		return nil
	}
	return db
}

func MongoDB() *mongo.Client {
	mongodbOptions := options.Client().ApplyURI("mongodb://localhost:27017/chat")
	mongodb, err := mongo.Connect(context.Background(), mongodbOptions)
	if err != nil {
		logging.Error("MysqlDB connect failed : ", err)
		return nil
	}
	return mongodb
}

func initDB() {
	global.Config.MysqlDB = MysqlDB()
	if global.Config.MysqlDB == nil {
		logging.Error("MysqlDB connect failed !")
	}
	//global.Config.MongoDB = MongoDB()
	//if global.Config.MongoDB == nil {
	//	return
	//}
}
