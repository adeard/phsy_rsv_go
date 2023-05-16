package config

import (
	"fmt"
	"log"
	"os"
	"phsy_rsv_go/modules/book"
	"phsy_rsv_go/modules/rate"
	"phsy_rsv_go/modules/user"
	userlevel "phsy_rsv_go/modules/user_level"
	usertype "phsy_rsv_go/modules/user_type"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func Connect() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	switch Dbdriver {
	case "sqlsvr":
		SqlsvrDev(DbUser, DbPassword, DbHost, DbName, Dbdriver)
	case "mysql":
		MysqlDev(DbUser, DbPassword, DbHost, DbPort, DbName, Dbdriver)
	}

	migrateDB()
	seedDB()

	return Db
}

func MysqlDev(DbUser string, DbPassword string, DbHost string, DbPort string, DbName string, Dbdriver string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("Database Connection Error")
	}

}

func SqlsvrDev(DbUser string, DbPassword string, DbHost string, DbName string, Dbdriver string) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=disable&connection+timeout=30", DbUser, DbPassword, DbHost, DbName)
	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("Database Connection Error")
	}
}

func migrateDB() {
	Db.AutoMigrate(&book.Book{})
	Db.AutoMigrate(&user.User{})
	Db.AutoMigrate(&usertype.UserType{})
	Db.AutoMigrate(&userlevel.UserLevel{})
	Db.AutoMigrate(&rate.Rate{})
}

func seedDB() {
	InsertUserLevel()
	InsertUserType()
}

func InsertUserLevel() {
	var total int64
	var ul userlevel.UserLevel
	Db.Model(&ul).Count(&total)
	if total == 0 {
		var data = []userlevel.UserLevel{
			{
				Name: "admin", IsActive: true,
			},
			{
				Name: "user", IsActive: true,
			},
		}

		Db.Create(&data)
	}
}

func InsertUserType() {
	var total int64
	var ut usertype.UserType
	Db.Model(&ut).Count(&total)
	if total == 0 {
		var data = []usertype.UserType{
			{
				Name: "patient", IsActive: true,
			},
			{
				Name: "staff", IsActive: true,
			},
		}

		Db.Create(&data)
	}
}
