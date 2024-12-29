package global

import (
	"GopherMall/user_srv/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

const (
	DBUser     = "root"
	DBPassword = "Outer233"
	DBAddr     = "127.0.0.1:3306"
	DBName     = "GopherMall"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBAddr,
		DBName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	_ = DB.AutoMigrate(&model.User{})
}
