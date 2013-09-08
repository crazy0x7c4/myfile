package service

import (
	"dao"
	. "entity"
	"log"
)

var folderDao = new(dao.FolderDao)

type FolderService struct {
}

func (folderService *FolderService) AddFolder(folder *Folder) (int, error) {
	id, err := folderDao.AddFolder(folder)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func (folderService *FolderService) DelFolder(id int) error {
	err := folderDao.DelFolder(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (folderService *FolderService) ModFolder(folder *Folder) error {
	err := folderDao.ModFolder(folder)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (folderService *FolderService) GetFolder(id int) (*Folder, error) {
	folder, err := folderDao.GetFolder(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return folder, nil
}

func (folderService *FolderService) GetFolderList() ([]*Folder, error) {
	folderList, err := folderDao.GetFolderList()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return folderList, nil
}
