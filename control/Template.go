package control

import (
	"html/template"
	"io/ioutil"
	"log"
	"path"
)

const TEMPLATE_DIR = "web/html"

var Templates = make(map[string]*template.Template)

func init() {
	fileInfos, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if ext := path.Ext(fileName); ext == ".html" {
			filePath := TEMPLATE_DIR + "/" + fileName
			t := template.Must(template.ParseFiles(filePath))
			Templates[fileName] = t
			log.Println("Loaded:" + fileName)
		}
	}

}
