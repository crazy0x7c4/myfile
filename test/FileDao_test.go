package test

//import (
//	"dao"
//	. "entity"
//	"log"
//	"testing"
//	"time"
//)

//func TestAddFile(t *testing.T) {
//	fileDao := new(dao.FileDao)
//	file := File{1, 1, 1, "test.txt", ".txt", 1024, time.Now().Unix(), time.Now().Unix()}
//	_, err := fileDao.AddFile(&file)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestDelFile(t *testing.T) {
//	fileDao := new(dao.FileDao)
//	fileDao.DelFile(2)
//}

//func TestModFile(t *testing.T) {
//	fileDao := new(dao.FileDao)
//	file, err := fileDao.GetFile(3)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	file.FolderId = 2
//	file.UserId = 3
//	file.Name = "mod.txt"
//	file.Stuffix = ".exe"
//	file.Size = 2048
//	file.CreateDate = time.Now().Unix()
//	file.ModifyDate = time.Now().Unix()
//	errMod := fileDao.ModFile(file)
//	if errMod != nil {
//		log.Println(errMod.Error())
//	}
//}

//func TestGetFile(t *testing.T) {
//	fileDao := new(dao.FileDao)
//	file, err := fileDao.GetFile(3)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	log.Println(file)
//}

//func TestGetFileList(t *testing.T) {
//	fileDao := new(dao.FileDao)
//	fileSlice, err := fileDao.GetFileList()
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for _, file := range fileSlice {
//		log.Println(file)
//	}
//}
