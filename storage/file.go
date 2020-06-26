package storage

import (
	"errors"
	"fmt"

	mydb "github.com/stsfang/go-netdisk/storage/mysql"
)

type TableFile struct {
	FileHash string
	FileName string
	FileSize int64
	FileAddr string
}

//文件元数据的增删改查CRUD

/**
* InsertFileMetaInfo 新增一条文件元数据记录
* @param ...
* @return 操作失败则false，成功则true
 */
func InsertFileMetaInfo(filehash, filename string, filesize int64, fileaddr string) bool {
	if mydb.DBConn() == nil {
		fmt.Println("jdhdhdhdhdh")
		return false
	}
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`," +
			" `file_addr`, `status`) values (?,?,?,?,?)")
	if err != nil || stmt == nil {
		fmt.Printf("获取statimemnt失败 %s\n", err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr, 1)
	if err != nil {
		fmt.Printf("stmt执行失败 %s \n", err.Error())
		return false
	}

	fmt.Println("dddd")
	//执行成功，检查是否有生效
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("stmt执行成功但0生效")
		}
		return true
	}

	return false
}

/**
* GetFileMeta 根据文件filehash查询文件元数据
* @param filhash 文件hash
* @return TableFile
 */
func GetFileMeta(filehash string) (*TableFile, error) {
	return &TableFile{}, errors.New("")
}

//
// GetFileMetaList 获取多条文件元数据
// @param limit 一次查询的个数
// @return
//
func GetFileMetaList(limit int) ([]TableFile, error) {
	return make([]TableFile, 0), errors.New("")
}

/**
* UpdateFileLocation 更新文件存储位置
* @param ...
* @return
 */
func UpdateFileLocation(filehash, fileaddr string) bool {
	return false
}
