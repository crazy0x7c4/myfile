package dao

import (
	. "entity"
	"strconv"
)

type FolderDao struct {
}

func (folderDao *FolderDao) AddFolder(folder *Folder) (int, error) {
	sqlStr := "INSERT INTO folder VALUES (?,?,?,?,?)"
	id, err := Execute(sqlStr, 0, folder.Id, folder.UserId, folder.Name, folder.CreateDate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (folderDao *FolderDao) DelFolder(id int) error {
	sqlStr := "DELETE FROM folder WHERE id=?"
	_, err := Execute(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (folderDao *FolderDao) ModFolder(folder *Folder) error {
	sqlStr := "UPDATE folder SET parent_id=?, user_id=?, name=?, create_date=? WHERE id=?"
	_, err := Execute(sqlStr, folder.ParentId, folder.UserId, folder.Name, folder.CreateDate, folder.Id)
	if err != nil {
		return err
	}
	return nil
}

func (folderDao *FolderDao) GetFolder(id int) (*Folder, error) {
	sqlStr := "SELECT * FROM folder WHERE id=?"
	data, err := QueryRow(5, sqlStr, id)
	if err != nil {
		return nil, err
	}
	folder := folderDao.setData(data)
	return folder, nil
}

func (folderDao *FolderDao) GetFolderList() ([]*Folder, error) {
	sqlStr := "SELECT * FROM folder"
	data, err := Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var folderList []*Folder
	for _, folderSlice := range data {
		folder := folderDao.setData(folderSlice)
		folderList = append(folderList, folder)
	}
	return folderList, nil
}

func (folderDao *FolderDao) setData(data []string) *Folder {
	folder := new(Folder)

	int32Id, _ := strconv.ParseInt(data[0], 0, 32)
	folder.Id = int(int32Id)

	int32ParentId, _ := strconv.ParseInt(data[1], 0, 32)
	folder.ParentId = int(int32ParentId)

	int32UserId, _ := strconv.ParseInt(data[2], 0, 32)
	folder.UserId = int(int32UserId)

	folder.Name = data[3]
	int64CreateDate, _ := strconv.ParseInt(data[4], 0, 32)

	folder.CreateDate = int64CreateDate

	return folder
}
