package dao

import (
	. "entity"
	"strconv"
)

type FileDao struct {
}

func (fileDao *FileDao) AddFile(file *File) (int, error) {
	sqlStr := "INSERT INTO file VALUES (?,?,?,?,?,?,?,?)"
	id, err := Execute(sqlStr, 0, file.FolderId, file.UserId, file.Name,
		file.Stuffix, file.Size, file.CreateDate, file.ModifyDate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (fileDao *FileDao) DelFile(id uint) error {
	sqlStr := "DELETE FROM file WHERE id=?"
	_, err := Execute(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (fileDao *FileDao) ModFile(file *File) error {
	sqlStr := "UPDATE file SET folder_id=?, user_id=?, name=?," +
		"stuffix=?, size=?, create_date=?, modify_date=? WHERE id=?"
	_, err := Execute(sqlStr, file.FolderId, file.UserId, file.Name,
		file.Stuffix, file.Size, file.CreateDate, file.ModifyDate, file.Id)
	if err != nil {
		return err
	}
	return nil
}

func (fileDao *FileDao) GetFile(id uint) (*File, error) {
	sqlStr := "SELECT * FROM file WHERE id=?"
	data, err := QueryRow(8, sqlStr, id)
	if err != nil {
		return nil, err
	}
	file := fileDao.setData(data)
	return file, nil
}

func (fileDao *FileDao) GetFileList() ([]*File, error) {
	sqlStr := "SELECT * FROM file"
	data, err := Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var fileList []*File
	for _, fileSlice := range data {
		file := fileDao.setData(fileSlice)
		fileList = append(fileList, file)
	}
	return fileList, nil
}

func (fileDao *FileDao) setData(data []string) *File {
	file := new(File)
	int64Id, _ := strconv.ParseInt(data[0], 0, 32)
	file.Id = int(int64Id)
	int64FolderId, _ := strconv.ParseInt(data[1], 0, 32)
	file.FolderId = int(int64FolderId)
	int64UserId, _ := strconv.ParseInt(data[2], 0, 32)
	file.UserId = int(int64UserId)
	file.Name = data[3]
	file.Stuffix = data[4]
	file.Size, _ = strconv.ParseInt(data[5], 0, 64)
	int64CreateDate, _ := strconv.ParseInt(data[6], 0, 64)
	file.CreateDate = int64CreateDate
	int64ModifyDate, _ := strconv.ParseInt(data[7], 0, 64)
	file.ModifyDate = int64ModifyDate
	return file
}
