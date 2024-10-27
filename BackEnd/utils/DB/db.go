package DB

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sushi/model"
)

type DB struct {
	DB  *gorm.DB
	log *logrus.Logger
}

func NewDB_MySQL(log *logrus.Logger, DBPath string) *DB {
	dsn := fmt.Sprintf(DBPath)
	var _db *gorm.DB
	var err error

	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connection error: " + err.Error())
	}

	sqlDB, _ := _db.DB()
	//err = _db.AutoMigrate(model.Toponym{})
	//if err != nil {
	//	return nil
	//}
	err = _db.AutoMigrate(model.Unit{})
	if err != nil {
		return nil
	}

	sqlDB.SetMaxOpenConns(100) //连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //最大允许的空闲连接数
	return &DB{
		log: log,
		DB:  _db,
	}
}
