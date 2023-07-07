package initiailizers

import (
	"log"
	"os"

	"github.com/techdenglei/eshop/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}
	return db
}

// 迁移数据库
func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.CatalogBrand{})
	db.AutoMigrate(&models.CatalogItem{})
	db.AutoMigrate(&models.CatalogType{})
}
