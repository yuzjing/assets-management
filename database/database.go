package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "asset_reader:asset_reader@tcp(192.168.1.140:3306)/gdzc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to database")

	// err = db.AutoMigrate(&models.Asset{}, &models.AssetChangeLog{})
	// if err != nil {
	// 	fmt.Println("Failed to migrate database:", err)
	// 	panic("Failed to migrate database")
	// }

	// fmt.Println("Database migrated")
	DB = db
}
