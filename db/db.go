package db

import (
	"fmt"
	// "strconv"

	"github.com/bayadev/bayahtmx/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connect to db
func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&models.Category{},&models.Product{}, &models.User{}, &models.Image{},  &models.Order{}, &models.OrderItem{}, &models.Review{}, &models.Auction{}, &models.Visit{}, &models.State{}, &models.City{}, &models.Faq{}, &models.Filter{}, &models.FilterOption{}, &models.Promoted{}, &models.Contact{}, &models.Blog{}, &models.Shop{})
	fmt.Println("Database Migrated")
}
