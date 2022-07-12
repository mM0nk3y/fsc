package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "zhangsan"
	passWord = "P@ssyoudontknow886"
	ip       = "xxx.xxx.xxx.xxx"
	port     = "3306"
	dbName   = "banner"
)

func ConnectMySQL() *sql.DB {
	path := strings.Join([]string{userName, ":", passWord, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(1000)
	DB.SetMaxIdleConns(100)

	if err := DB.Ping(); err != nil {
		panic(err)
		fmt.Println("MySQL Connect failed")
	}
	fmt.Println("MySQL Connect success")
	return DB

}

func Quchong() {
	fmt.Println("quchong test")
}
