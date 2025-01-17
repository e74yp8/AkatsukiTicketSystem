package sql

import (
	"TicketBackend/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSQL() {
	cfg := config.LoadConfig()
	username := cfg.MySQL.Username
	password := cfg.MySQL.Password
	host := cfg.MySQL.Host
	port := cfg.MySQL.Port
	Dbname := cfg.MySQL.Database
	timeout := "10s"

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Connect Database Failed, Error=" + err.Error())
	}

	//Connection succeed
	//fmt.Println(db)

	DB = db
}
