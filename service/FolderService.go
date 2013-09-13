package service

import (
	. "dao"
	"log"
	"os"
	"strconv"
)

var folderDao = new(Folder)

type FolderService struct {
}

func (this *FolderService) Create(folder *Folder) (int, error) {
	id, err := folderDao.Add(folder)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func (this *FolderService) Delete(id int) error {
	//删除文件夹
	condition := make(map[string]interface{})
	condition["f_id"] = id
	errDel := folderDao.Delete(condition)
	if errDel != nil {
		log.Println(errDel.Error())
		return errDel
	}

	//删除文件夹里的文件
	condition = make(map[string]interface{})
	condition["f_folder_id"] = id
	data, errGetOne := fileDao.GetList(condition)
	if errGetOne != nil {
		log.Println(errGetOne.Error())
		return errGetOne
	}
	for _, v := range data {
		fileIdStr := strconv.Itoa(v.Id)
		_, err := os.Stat(UPLOAD_DIR + "/" + fileIdStr + "/" + v.Name)
		if err != nil {
			return err
		} else {
			errRemoveFile := os.Remove(UPLOAD_DIR + "/" + fileIdStr + "/" + v.Name)
			if errRemoveFile != nil {
				log.Println(errRemoveFile.Error())
				return errRemoveFile
			}
			errRemoveFolder := os.Remove(UPLOAD_DIR + "/" + fileIdStr)
			if errRemoveFolder != nil {
				log.Println(errRemoveFolder.Error())
				return errRemoveFolder
			}
		}
		condition = make(map[string]interface{})
		condition["f_id"] = v.Id
		errDelete := fileDao.Delete(condition)
		if errDelete != nil {
			log.Println(errDelete.Error())
			return errDelete
		}
	}

	//获取子文件夹，递归删除
	condition = make(map[string]interface{})
	condition["f_pid"] = id
	subfolder, errGetList := folderDao.GetList(condition)
	if errGetList != nil {
		log.Println(errGetList.Error())
		return errGetList
	}

	for _, v := range subfolder {
		return this.Delete(v.Id)
	}

	return nil
}

func (this *FolderService) Rename(id int, name string) error {
	columns := make(map[string]interface{})
	columns["f_name"] = name
	condition := make(map[string]interface{})
	condition["f_id"] = id
	err := folderDao.Update(columns, condition)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (this *FolderService) GetRoot(accountId int) ([]*Folder, error) {
	condition := make(map[string]interface{})
	condition["f_account_id"] = accountId
	condition["f_pid"] = 0
	data, err := folderDao.GetList(condition)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return data, nil
}

func (this *FolderService) GetSubfolder(parentId int) ([]*Folder, error) {
	condition := make(map[string]interface{})
	condition["f_pid"] = parentId
	data, err := folderDao.GetList(condition)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return data, nil
}

func (this *FolderService) GetPath(id int) ([]int, []string, error) {
	var pathStr []string
	var pathInt []int

	pathStr = append(pathStr, "Home")
	pathInt = append(pathInt, 0)

	columns := []string{"f_name", "f_pid"}
	condition := make(map[string]interface{})
	condition["f_id"] = id

	nodeStr, err := folderDao.GetOneColumn(columns, condition)
	if err != nil {
		return nil, nil, err
	}

	pathStr = append(pathStr, nodeStr[0])
	pathInt = append(pathInt, id)

	pid, _ := strconv.Atoi(nodeStr[1])
	for pid != 0 {
		condition := make(map[string]interface{})
		condition["f_id"] = pid
		nodeStr, err := folderDao.GetOneColumn(columns, condition)
		if err != nil {
			return nil, nil, err
		}
		pathStr = append(pathStr, nodeStr[0])
		pathInt = append(pathInt, pid)

		pid, _ = strconv.Atoi(nodeStr[1])
	}

	var reverseInt []int
	var reverseStr []string

	for i := len(pathInt); i > 0; i-- {
		reverseInt = append(reverseInt, pathInt[i-1])
		reverseStr = append(reverseStr, pathStr[i-1])
	}
	return reverseInt, reverseStr, nil
}
