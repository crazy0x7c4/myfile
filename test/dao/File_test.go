package test

//import (
//	"dao"
//	"log"
//	"testing"
//	"time"
//)

//func TestAdd(t *testing.T) {
//	fileDao := new(dao.File)
//	fileObj := dao.File{0, 10, 104, "cod2.txt", ".txt", 1024, time.Now().Unix(), time.Now().Unix()}
//	_, err := fileDao.Add(&fileObj)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestDelete(t *testing.T) {
//	fileDao := new(dao.File)
//	fileDao.Delete(3)
//}

//func TestModFile(t *testing.T) {
//	fileDao := new(dao.File)
//	fileObj, err := fileDao.GetById(4)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	fileObj.FolderId = 2
//	fileObj.UserId = 3
//	fileObj.Name = "mod.txt"
//	fileObj.Stuffix = ".exe"
//	fileObj.Size = 2048
//	fileObj.CreateDate = time.Now().Unix()
//	fileObj.ModifyDate = time.Now().Unix()
//	errMod := fileDao.ModFile(fileObj)
//	if errMod != nil {
//		log.Println(errMod.Error())
//	}
//}

//func TestGetById(t *testing.T) {
//	fileDao := new(dao.File)
//	fileObj, err := fileDao.GetById(4)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	log.Println(fileObj)
//}

//func TestGetByFolderId(t *testing.T) {
//	fileDao := new(dao.File)
//	fileSlice, err := fileDao.GetByFolderId(1)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for _, v := range fileSlice {
//		log.Println(v)
//	}
//}
