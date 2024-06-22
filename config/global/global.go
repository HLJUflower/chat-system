package global

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
)

type Global struct {
	Logger  *logrus.Logger
	MysqlDB *gorm.DB
	MongoDB *mongo.Client
}

var (
	Config = new(Global)
	WG     = sync.WaitGroup{}
)
