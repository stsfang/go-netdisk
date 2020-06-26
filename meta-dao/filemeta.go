package meta

import (
	"fmt"

	"github.com/stsfang/go-netdisk/storage"
)

//FileMeta 文件元数据
type FileMeta struct {
	FileSha1 string
	FileName string
	Location string
	UploadAt string
	FileSize int64
}

var fileMetaMap map[string]FileMeta

func init() {
	fileMetaMap = make(map[string]FileMeta)
}

//充当DAO层，提供数据操作接口

// AddFileMeta 新增一个文件元数据
func AddFileMeta(fm FileMeta) bool {
	return storage.InsertFileMetaInfo(fm.FileSha1, fm.FileName, fm.FileSize, fm.Location)
}

//UpdateFileMeta 更新文件元数据
func UpdateFileMeta(fm FileMeta) bool {
	return storage.UpdateFileLocation(fm.FileSha1, fm.Location)
}

//GetFileMeta 获取指定filesha1的文件元数据
func GetFileMeta(filesha1 string) (*FileMeta, error) {
	tableFile, err := storage.GetFileMeta(filesha1)
	if err != nil || tableFile == nil {
		return nil, err
	}
	var fileMeta = FileMeta{
		FileSha1: tableFile.FileHash,
		FileName: tableFile.FileName,
		FileSize: tableFile.FileSize,
		Location: tableFile.FileAddr,
	}

	fmt.Println(fileMeta)
	return &fileMeta, nil
}

//GetLastFileMetaList 获取最近limit条文件元数据
func GetLastFileMetaList(limit int) ([]FileMeta, error) {
	tFiles, err := storage.GetFileMetaList(limit)
	if err != nil {
		return make([]FileMeta, 0), err
	}

	resultFiles := make([]FileMeta, len(tFiles))
	for i := 0; i < len(tFiles); i++ {
		resultFiles[i] = FileMeta{
			FileSha1: tFiles[i].FileHash,
			FileName: tFiles[i].FileName,
			FileSize: tFiles[i].FileSize,
			Location: tFiles[i].FileAddr,
		}
	}
	return resultFiles, nil
}
