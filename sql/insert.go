package sql

import (
	"database/sql"
	"fmt"
	"time"
)

func InsertDB(db *sql.DB /*Id int, Task_ID int,*/, Target string, Banner string /*Server string, Status_Code int,*/, Title string /*, Last_Time int, Pocmatch int*/) {
	//db := ConnectMySQL()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	sql := "insert ignore into bigtask (`Id`,`Task_Id`,`Target`,`Banner`,`Server`,`Status_Code`,`Title`,`Last_Time`,`Pocmatch`)values(?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	content, err := stmt.Exec(nil, "", Target, Banner, "", "", Title, time.Now(), 0)
	if err != nil {
		panic(err)
	}
	tx.Commit()
	fmt.Println(content.LastInsertId())
	fmt.Println(content.RowsAffected())
}
