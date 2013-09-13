package test

//import (
//	"dao"
//	"log"
//	"testing"
//	"time"
//)

//func TestAdd(t *testing.T) {
//	folderDao := new(dao.Folder)

//	folderObj := dao.Folder{0, 24, 104, "ss5", time.Now().Unix()}
//	_, err := folderDao.Add(&folderObj)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestDelete(t *testing.T) {
//	folderDao := new(dao.Folder)
//	err := folderDao.Delete(3)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestModName(t *testing.T) {
//	folderDao := new(dao.Folder)
//	err := folderDao.ModName(7, "Only Mod Name")
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestModFolder(t *testing.T) {
//	folderDao := new(dao.Folder)
//	folderObj, err := folderDao.GetById(7)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	folderObj.ParentId = 2
//	folderObj.UserId = 105
//	folderObj.Name = "ModName"
//	folderObj.CreateDate = time.Now().Unix()
//	errMod := folderDao.ModFolder(folderObj)
//	if errMod != nil {
//		log.Println(errMod.Error())
//	}
//}

//func TestGetRootByUserId(t *testing.T) {
//	folderDao := new(dao.Folder)
//	data, err := folderDao.GetRootByUserId(104)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for _, v := range data {
//		log.Println(v.Name)
//	}
//}

//func TestGetByParentId(t *testing.T) {
//	folderDao := new(dao.Folder)
//	data, err := folderDao.GetByParentId(2)
//	if err != nil {
//		log.Println(err)
//	}
//	for _, v := range data {
//		log.Println(v)
//	}
//}
