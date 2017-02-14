package models

import (
	"fmt"
	"gin-template/config"
	"log"

	"github.com/jinzhu/gorm"
)

// DB db instance
var (
	DB    *gorm.DB
	DBErr error
)

// BaseModel is gorm base model template
type BaseModel struct {
	gorm.Model
}

// NewRecord return object is new record
func (m BaseModel) NewRecord() bool {
	return m.ID <= 0
}

// Create return model create status
func (m BaseModel) Create() error {
	return DB.Create(&m).Error
}

// Save return model save status
func (m BaseModel) Save() error {
	return DB.Save(&m).Error
}

// Destroy object from database
func (m BaseModel) Destroy() error {
	return DB.Delete(&m).Error
}

// IsDeleted return object is deleted
func (m BaseModel) IsDeleted() bool {
	return m.DeletedAt != nil
}

//InitDB initialize database connection
func InitDB() {
	connectDatabase()
	DB.LogMode(true)
	DB.SetLogger(log.New(config.LogFile(), "\n", 3))
	migrateTables()
}

func migrateTables() {
	DB.AutoMigrate(&User{})
}

func connectDatabase() {
	var adapter, option string
	dbcfg := config.GetSection("database")
	adapter = dbcfg.Key("adapter").Value()
	hostname := dbcfg.Key("hostname").Value()
	username := dbcfg.Key("username").Value()
	password := dbcfg.Key("password").Value()
	database := dbcfg.Key("database").Value()
	port := dbcfg.Key("port").Value()

	if port == "" {
		switch adapter {
		case "postgres":
			port = "5432"
		case "mysql":
			port = "3306"
		}
	}

	switch adapter {
	case "postgres":
		option = fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", hostname, username, password, database)
	case "mysql":
		option = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", username, password, hostname, port, database)
	case "sqlite3", "sqlite":
		option = fmt.Sprintf("db/%v.db", database)
	}
	fmt.Println(option)

	DB, DBErr = gorm.Open(adapter, option)
	if DBErr != nil {
		fmt.Println("connection database error!")
		fmt.Println(DBErr)
		panic(DBErr)
	}
}
