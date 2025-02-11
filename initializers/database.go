package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	mysql_host := os.Getenv("MYSQL_HOST")
	mysql_port := os.Getenv("MYSQL_PORT")
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASSWORD")
	mysql_database := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql_user, mysql_password, mysql_host, mysql_port, mysql_database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}