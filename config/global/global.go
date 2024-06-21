package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type Global struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

var (
	Config = new(Global)
	WG     = sync.WaitGroup{}
)
