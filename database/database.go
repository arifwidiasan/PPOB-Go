package database

import (
	"fmt"
	"time"

	"github.com/CapstoneProject31/backend_ppob_31/config"
	"github.com/CapstoneProject31/backend_ppob_31/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	DB, err := gorm.Open(mysql.Open(conectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	admin := DB.Migrator().HasTable(&model.Admin{})
	if !admin {
		DB.Migrator().CreateTable(&model.Admin{})
		DB.Model(&model.Admin{}).Create([]map[string]interface{}{
			{"username": "admin", "password": "admin123", "created_at": time.Now()},
		})
	}
	DB.AutoMigrate(&model.Product_type{}, &model.Operator{})
	return DB
}
