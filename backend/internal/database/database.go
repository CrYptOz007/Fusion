package database

import (
	"fmt"

	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/CrYptOz007/Fusion/internal/models/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connector interface {
	Init()
}

type Connection struct {
	Db *gorm.DB
}

func (c *Connection) Init(error chan<- error) {
	fmt.Println("Initializing database connection")

	dsn := "root@tcp(backend_db:3306)/fusion?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		error <- err
	}

	c.Db = db

	c.autoMigration()

	fmt.Println("Database connected")
	error <- nil
}

func (c *Connection) autoMigration() {
	c.Db.AutoMigrate(&service.Service{})
	c.Db.AutoMigrate(&user.User{})
}
