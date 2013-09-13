package dao

import (
	"strconv"
)

type File struct {
	Id         int
	FolderId   int
	AccountId  int
	Name       string
	Stuffix    string
	Size       int64
	CreateDate int64
	ModifyDate int64
}

func (this *File) Add(file *File) (int, error) {
	id, err := Add("file", 0, file.FolderId, file.AccountId, file.Name,
		file.Stuffix, file.Size, file.CreateDate, file.ModifyDate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *File) Delete(condition map[string]interface{}) error {
	err := Delete("file", condition)
	if err != nil {
		return err
	}
	return nil
}

func (this *File) Update(columns map[string]interface{}, condition map[string]interface{}) error {
	err := Update("file", columns, condition)
	if err != nil {
		return err
	}
	return nil
}

func (this *File) GetList(condition map[string]interface{}) ([]*File, error) {
	data, err := GetList("file", condition)
	if err != nil {
		return nil, err
	}
	list := this.parseDataList(data)
	return list, nil
}

func (this *File) GetOne(condition map[string]interface{}) (*File, error) {
	data, err := GetOne("file", 8, condition)
	if err != nil {
		return nil, err
	}
	obj := this.parseData(data)
	return obj, nil
}

func (this *File) parseData(data []string) *File {
	obj := new(File)

	id, _ := strconv.Atoi(data[0])
	obj.Id = id

	folderId, _ := strconv.Atoi(data[1])
	obj.FolderId = folderId

	accountId, _ := strconv.Atoi(data[2])
	obj.AccountId = accountId

	obj.Name = data[3]
	obj.Stuffix = data[4]

	obj.Size, _ = strconv.ParseInt(data[5], 0, 64)

	createDate, _ := strconv.ParseInt(data[6], 0, 64)
	obj.CreateDate = createDate

	modifyDate, _ := strconv.ParseInt(data[7], 0, 64)
	obj.ModifyDate = modifyDate

	return obj
}

func (this *File) parseDataList(data [][]string) []*File {
	var list []*File
	for i := 0; i < len(data); i++ {
		obj := this.parseData(data[i])
		list = append(list, obj)
	}
	return list
}

//func (this *File) ModName(id int, name string) error {
//	sqlStr := "UPDATE file SET name=? WHERE id=?"
//	_, err := Execute(sqlStr, name, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (this *File) ModModifyDate(id int, date int64) error {
//	sqlStr := "UPDATE file SET modify_Date=? WHERE id=?"
//	_, err := Execute(sqlStr, date, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (this *File) ModSize(id int, size int64) error {
//	sqlStr := "UPDATE file SET size=? WHERE id=?"
//	_, err := Execute(sqlStr, size, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (this *File) ModFile(p *File) error {
//	sqlStr := "UPDATE file SET folder_id=?, user_id=?, name=?," +
//		"stuffix=?, size=?, create_date=?, modify_date=? WHERE id=?"
//	_, err := Execute(sqlStr, p.FolderId, p.UserId, p.Name,
//		p.Stuffix, p.Size, p.CreateDate, p.ModifyDate, p.Id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (this *File) GetById(id int) (*File, error) {
//	sqlStr := "SELECT * FROM file WHERE id=?"
//	data, err := QueryRow(8, sqlStr, id)
//	if err != nil {
//		return nil, err
//	}
//	p := this.setData(data)
//	return p, nil
//}

//func (this *File) GetRootByUserId(userId int) ([]*File, error) {
//	sqlStr := "SELECT * FROM file WHERE user_id=? AND folder_id=0"
//	data, err := Query(sqlStr, userId)
//	if err != nil {
//		return nil, err
//	}
//	var fileList []*File
//	for _, fileSlice := range data {
//		p := this.setData(fileSlice)
//		fileList = append(fileList, p)
//	}
//	return fileList, nil
//}

//func (this *File) GetByFolderId(folderId int) ([]*File, error) {
//	sqlStr := "SELECT * FROM file WHERE folder_id=?"
//	data, err := Query(sqlStr, folderId)
//	if err != nil {
//		return nil, err
//	}
//	var fileList []*File
//	for _, fileSlice := range data {
//		p := this.setData(fileSlice)
//		fileList = append(fileList, p)
//	}
//	return fileList, nil
//}

//func (this *File) GetByFolderIdName(userId int, folderId int, name string) (*File, error) {
//	sqlStr := "SELECT * FROM file WHERE user_id=? AND folder_id=? AND name=?"
//	data, err := QueryRow(8, sqlStr, userId, folderId, name)
//	if err != nil {
//		return nil, err
//	}
//	p := this.setData(data)
//	return p, nil
//}
