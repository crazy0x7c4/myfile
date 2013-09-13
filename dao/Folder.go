package dao

import (
	"strconv"
)

type Folder struct {
	Id         int
	Pid        int
	AccountId  int
	Name       string
	CreateDate int64
}

func (this *Folder) Add(folder *Folder) (int, error) {
	id, err := Add("folder", 0, folder.Pid, folder.AccountId, folder.Name, folder.CreateDate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *Folder) Delete(condition map[string]interface{}) error {
	err := Delete("folder", condition)
	if err != nil {
		return err
	}
	return nil
}

func (this *Folder) Update(columns map[string]interface{}, condition map[string]interface{}) error {
	err := Update("folder", columns, condition)
	if err != nil {
		return err
	}
	return nil
}

func (this *Folder) GetList(condition map[string]interface{}) ([]*Folder, error) {
	data, err := GetList("folder", condition)
	if err != nil {
		return nil, err
	}

	list := this.parseDataList(data)

	return list, nil
}

func (this *Folder) GetOne(condition map[string]interface{}) (*Folder, error) {
	data, err := GetOne("folder", 5, condition)
	if err != nil {
		return nil, err
	}
	obj := this.parseData(data)
	return obj, nil
}

func (this *Folder) GetOneColumn(columns []string, condition map[string]interface{}) ([]string, error) {
	data, err := GetOneColumn("folder", columns, condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *Folder) parseData(data []string) *Folder {
	obj := new(Folder)

	id, _ := strconv.Atoi(data[0])
	obj.Id = id

	pid, _ := strconv.Atoi(data[1])
	obj.Pid = pid

	accountId, _ := strconv.Atoi(data[2])
	obj.AccountId = accountId

	obj.Name = data[3]

	createDate, _ := strconv.ParseInt(data[4], 0, 32)
	obj.CreateDate = createDate

	return obj
}

func (this *Folder) parseDataList(data [][]string) []*Folder {
	var list []*Folder
	for i := 0; i < len(data); i++ {
		obj := this.parseData(data[i])
		list = append(list, obj)
	}
	return list
}

//func (folder *Folder) ModName(id int, name string) error {
//	sqlStr := "UPDATE folder SET name=? WHERE id=?"
//	_, err := Execute(sqlStr, name, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (folder *Folder) ModFolder(p *Folder) error {
//	sqlStr := "UPDATE folder SET parent_id=?, user_id=?, name=?, create_date=? WHERE id=?"
//	_, err := Execute(sqlStr, p.ParentId, p.UserId, p.Name, p.CreateDate, p.Id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (folder *Folder) GetById(id int) (*Folder, error) {
//	sqlStr := "SELECT * FROM folder WHERE id=?"
//	data, err := QueryRow(5, sqlStr, id)
//	if err != nil {
//		return nil, err
//	}
//	p := folder.setData(data)
//	return p, nil
//}

//func (folder *Folder) GetRootByUserId(userId int) ([]*Folder, error) {
//	sqlStr := "SELECT * FROM folder WHERE user_id=? AND parent_id=0"
//	data, err := Query(sqlStr, userId)
//	if err != nil {
//		return nil, err
//	}
//	var folderList []*Folder
//	for _, folderSlice := range data {
//		p := folder.setData(folderSlice)
//		folderList = append(folderList, p)
//	}
//	return folderList, nil
//}

//func (folder *Folder) GetByParentId(parentId int) ([]*Folder, error) {
//	sqlStr := "SELECT * FROM folder WHERE parent_id=?"
//	data, err := Query(sqlStr, parentId)
//	if err != nil {
//		return nil, err
//	}
//	var folderList []*Folder
//	for _, folderSlice := range data {
//		p := folder.setData(folderSlice)
//		folderList = append(folderList, p)
//	}
//	return folderList, nil
//}

//func (folder *Folder) GetParentIdById(id int) (int, error) {
//	sqlStr := "SELECT parent_id FROM folder WHERE id=?"
//	data, err := QueryRow(1, sqlStr, id)
//	if err != nil {
//		return 0, err
//	}
//	parentId, errConv := strconv.Atoi(data[0])
//	if errConv != nil {
//		return 0, errConv
//	}
//	return parentId, nil
//}

//func (folder *Folder) GetName(id int) (string, error) {
//	sqlStr := "SELECT name FROM folder WHERE id=?"
//	data, err := QueryRow(1, sqlStr, id)
//	if err != nil {
//		return "", err
//	}
//	return data[0], nil
//}
