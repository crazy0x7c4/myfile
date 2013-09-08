package test

//import (
//	"dao"
//	. "entity"
//	"log"
//	"testing"
//	"time"
//)

//func TestAddFolder(t *testing.T) {
//	folderDao := new(dao.FolderDao)

//	folder := Folder{1, 0, 104, "f1", time.Now().Unix()}
//	_, err := folderDao.AddFolder(&folder)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestDelFolder(t *testing.T) {
//	folderDao := new(dao.FolderDao)
//	err := folderDao.DelFolder(1)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}

//func TestModFolder(t *testing.T) {
//	folderDao := new(dao.FolderDao)
//	folder, err := folderDao.GetFolder(2)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	folder.ParentId = 2
//	folder.UserId = 105
//	folder.Name = "ModName"
//	folder.CreateDate = time.Now().Unix()
//	errMod := folderDao.ModFolder(folder)
//	if errMod != nil {
//		log.Println(errMod.Error())
//	}
//}

//func TestGetFolder(t *testing.T) {
//	folderDao := new(dao.FolderDao)
//	folder, err := folderDao.GetFolder(2)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	log.Println(folder)
//}

//func TestGetFolderList(t *testing.T) {
//	folderDao := new(dao.FolderDao)
//	data, err := folderDao.GetFolderList()
//	if err != nil {
//		log.Println(err)
//	}
//	for _, v := range data {
//		log.Println(v)
//	}
//}
