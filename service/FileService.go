package service

import (
	. "dao"
	"io"
	"log"
	"os"
	"strconv"
)

const UPLOAD_DIR = "web/uploads"

var fileDao = new(File)

type FileService struct {
}

func (this *FileService) Upload(file *File, src io.Reader) error {
	condition := make(map[string]interface{})
	condition["f_account_id"] = file.AccountId
	condition["f_folder_id"] = file.FolderId
	condition["f_name"] = file.Name
	oldFile, err := fileDao.GetOne(condition)
	if err != nil && oldFile == nil {
		id, errAdd := fileDao.Add(file)
		if errAdd != nil {
			log.Println(errAdd.Error())
			return errAdd
		}
		file.Id = id
	} else {
		log.Println("文件已经存在，覆盖原文件")
		columns := make(map[string]interface{})
		columns["f_modify_date"] = file.ModifyDate
		condition := make(map[string]interface{})
		condition["f_id"] = oldFile.Id
		errUpdate := fileDao.Update(columns, condition)
		if errUpdate != nil {
			log.Println(errUpdate.Error())
			return errUpdate
		}
		file.Id = oldFile.Id
	}

	//上传或覆盖原文件
	fileIdStr := strconv.Itoa(file.Id)
	os.Mkdir(UPLOAD_DIR+"/"+fileIdStr, 0777)
	dst, errCreate := os.Create(UPLOAD_DIR + "/" + fileIdStr + "/" + file.Name)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return errCreate
	}
	defer dst.Close()

	size, errCopy := io.Copy(dst, src)
	if errCopy != nil {
		log.Println(errCopy.Error())
		return errCopy
	}

	columns := make(map[string]interface{})
	columns["f_size"] = size
	condition = make(map[string]interface{})
	condition["f_id"] = file.Id
	errModSize := fileDao.Update(columns, condition)
	if errModSize != nil {
		log.Println(errCopy.Error())
		return errCopy
	}

	return nil
}

func (this *FileService) Delete(id int) error {
	condition := make(map[string]interface{})
	condition["f_id"] = id
	file, errGetOne := fileDao.GetOne(condition)
	if errGetOne != nil {
		log.Println(errGetOne.Error())
		return errGetOne
	}
	fileIdStr := strconv.Itoa(file.Id)
	_, err := os.Stat(UPLOAD_DIR + "/" + fileIdStr + "/" + file.Name)
	if err != nil {
		return err
	} else {
		if errRemove := os.Remove(UPLOAD_DIR + "/" + fileIdStr + "/" + file.Name); errRemove != nil {
			log.Println(errRemove.Error())
			return errRemove
		}
		if errRemove := os.Remove(UPLOAD_DIR + "/" + fileIdStr); errRemove != nil {
			log.Println(errRemove.Error())
			return errRemove
		}
	}

	errDelFile := fileDao.Delete(condition)
	if errDelFile != nil {
		log.Println(errDelFile.Error())
		return errDelFile
	}

	return nil
}

func (this *FileService) Rename(id int, name string) error {
	columns := make(map[string]interface{})
	columns["f_name"] = name
	condition := make(map[string]interface{})
	condition["f_id"] = id
	err := fileDao.Update(columns, condition)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (this *FileService) GetRoot(accountId int) ([]*File, error) {
	condition := make(map[string]interface{})
	condition["f_account_id"] = accountId
	condition["f_folder_id"] = 0
	data, err := fileDao.GetList(condition)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *FileService) GetByFolderId(folderId int) ([]*File, error) {
	condition := make(map[string]interface{})
	condition["f_folder_id"] = folderId
	data, err := fileDao.GetList(condition)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return data, nil
}
