package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//需要导入mysql包，使用该包的init方法初始化mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//初始化 mysql db 连接
func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/netdisk?charset=utf8")
	if db == nil {
		fmt.Println("open mysql fail")
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(100)
	err := db.Ping()
	if err != nil {
		fmt.Printf("无法连接mysql，请检查 %s\n", err.Error())
		os.Exit(1)
	}
}

//DBConn 对外暴露mysql DB 连接
func DBConn() *sql.DB {
	return db
}

//RecordRow 类型别名
type RecordRow = map[string]interface{}

//ParseRows 从sql查询结果中解析出每一行，以map的形式记录每一列的key-value
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	//colums表的每一列
	colums, _ := rows.Columns()
	scanArgs := make([]interface{}, len(colums))
	values := make([]interface{}, len(colums))

	//将values的指针保存到scanArgs
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(RecordRow)
	records := make([]RecordRow, 0)

	//遍历每一行
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		//将扫描的行，记录到record
		for i, colValue := range values {
			if colValue != nil {
				record[colums[i]] = colValue
			}
		}

		records = append(records, record)
	}
	return records
}