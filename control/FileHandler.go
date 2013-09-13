package control

import (
	. "dao"
	"log"
	"net/http"
	"service"
	"strconv"
	"strings"
	"time"
)

var folderService = new(service.FolderService)
var fileService = new(service.FileService)

//文件列表
func ListHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	idStr := r.FormValue("id")
	id := atoi(idStr)
	data["id"] = id

	if id == 0 {
		accountId := getAccountId(r)
		//获取根文件夹
		folders, err := folderService.GetRoot(accountId)
		checkError(err)
		data["folders"] = folders

		//获取根文件
		files, err := fileService.GetRoot(accountId)
		checkError(err)
		data["files"] = files

		//路径
		path := make(map[string]int)
		var pathStr []string
		var pathInt []int

		pathStr = append(pathStr, "Home")
		pathInt = append(pathInt, 0)

		path[pathStr[0]] = pathInt[0]
		data["path"] = path
	} else {
		//获取子文件夹
		folders, err := folderService.GetSubfolder(id)
		checkError(err)
		data["folders"] = folders
		//获取子文件夹的文件
		files, err := fileService.GetByFolderId(id)
		checkError(err)
		data["files"] = files
		//获取路径
		pathInt, pathStr, err := folderService.GetPath(id)
		checkError(err)
		path := make(map[string]int)
		for i := 0; i < len(pathInt); i++ {
			path[pathStr[i]] = pathInt[i]
		}
		data["path"] = path
	}

	Templates["list.html"].Execute(w, data)
}

//创建文件夹
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	accountId := getAccountId(r)

	idStr := r.FormValue("id")
	id := atoi(idStr)

	name := r.FormValue("name")

	folder := new(Folder)
	folder.Pid = id
	folder.AccountId = accountId
	folder.Name = name
	folder.CreateDate = time.Now().Unix()
	_, errCreate := folderService.Create(folder)
	checkError(errCreate)

	http.Redirect(w, r, "/list?id="+idStr, http.StatusFound)
}

//删除文件夹
func DelFolderHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	folderId := atoi(idStr)
	err := folderService.Delete(folderId)
	checkError(err)
	pidStr := r.FormValue("pid")
	http.Redirect(w, r, "/list?id="+pidStr, http.StatusFound)
}

//上传文件
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	srcFile, fileHeader, err := r.FormFile("file")
	checkError(err)
	defer srcFile.Close()

	accountId := getAccountId(r)
	idStr := r.FormValue("id")
	id := atoi(idStr)

	file := new(File)
	file.FolderId = id
	file.AccountId = accountId
	file.Name = fileHeader.Filename
	lastIndex := strings.LastIndex(fileHeader.Filename, ".")
	file.Stuffix = fileHeader.Filename[lastIndex:]
	file.CreateDate = time.Now().Unix()
	file.ModifyDate = time.Now().Unix()
	errUpload := fileService.Upload(file, srcFile)
	checkError(errUpload)

	http.Redirect(w, r, "/list?id="+idStr, http.StatusFound)
}

//下载文件
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fileIdStr := r.FormValue("fileid")
	if fileIdStr == "" {
		fileIdStr = "0"
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, service.UPLOAD_DIR+"/"+fileIdStr+"/"+name)
}

//删除文件
func DelFileHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	fileId := atoi(idStr)
	err := fileService.Delete(fileId)
	checkError(err)
	folderidStr := r.FormValue("folderid")
	http.Redirect(w, r, "/list?id="+folderidStr, http.StatusFound)
}

//从cookie获取accountid
func getAccountId(r *http.Request) int {
	accountIdCookie, err := r.Cookie("AccountId")
	checkError(err)
	accountId := atoi(accountIdCookie.Value)
	return accountId
}

func atoi(str string) int {
	var i int
	if str == "" {
		i = 0
	} else {
		conv, err := strconv.Atoi(str)
		checkError(err)
		i = conv
	}
	return i
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

//func GeneratePath(id string) string {
//	if id == "0" {
//		return ""
//	}

//	folderId, err := strconv.Atoi(id)
//	checkError(err)
//	pathInt, pathStr, err := folderService.GetPath(folderId)
//	checkError(err)

//	var path string
//	for i := 0; i < len(pathInt); i++ {
//		temp := strconv.Itoa(pathInt[i])
//		path = "<a href=\"list?id=" + temp + "\">" + pathStr[i] + "</a>\\"
//	}
//	return path
//}
