package test

//import (
//	"dao"
//	"log"
//	"testing"
//)

//func TestExecute(t *testing.T) {
//	sqlStr := "INSERT account SET a_name=?, a_password=?"
//	id, _ := dao.Execute(sqlStr, "user1", "12345")
//	sqlStr = "DELETE FROM account WHERE a_id=?"
//	dao.Execute(sqlStr, id)
//}

//func TestQuery(t *testing.T) {
//	data, _ := dao.Query("SELECT * FROM account")
//	for _, v := range data {
//		log.Println(v)
//	}
//}

//func TestQueryRow(t *testing.T) {
//	data, _ := dao.QueryRow(3, "SELECT * FROM account WHERE a_name=? AND a_password=?", "crazy", "11111")
//	log.Println(data)
//}

//func TestAdd(t *testing.T) {
//	values := make([]interface{}, 8)
//	values[0] = 0
//	values[1] = 5
//	values[2] = 6
//	values[3] = "file5.txt"
//	values[4] = ".txt"
//	values[5] = 11304
//	values[6] = 1378972358
//	values[7] = 1378972358
//	id, err := dao.Add("file", values...)

//	//id, err := dao.Add("file", 0, 10, 6, "file6.txt", ".txt", 1024, 1378972358, 1378972358)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	log.Println(id)
//}

//func TestDelete(t *testing.T) {
//	condition := make(map[string]interface{})
//	condition["f_folder_id"] = "2"
//	condition["f_user_id"] = "3"
//	condition["f_name"] = "file2.txt"
//	err := dao.Delete("file", condition)
//	checkError(err)
//}

//func TestUpdate(t *testing.T) {
//	columns := make(map[string]interface{})
//	columns["f_folder_id"] = "13"
//	columns["f_user_id"] = "14"
//	columns["f_name"] = "filex.txt"

//	condition := make(map[string]interface{})
//	condition["f_id"] = "16"
//	condition["f_folder_id"] = "3"
//	condition["f_name"] = "file3.txt"
//	dao.Update("file", columns, condition)
//}

//func TestGetList(t *testing.T) {
//	condition := make(map[string]interface{})
//	condition["f_id"] = "16"
//	condition["f_folder_id"] = "13"
//	condition["f_name"] = "filex.txt"
//	data, err := dao.GetList("file", condition)
//	checkError(err)
//	for _, v := range data {
//		log.Println(v)
//	}
//}

//func TestGetOne(t *testing.T) {
//	condition := make(map[string]interface{})
//	condition["f_id"] = "16"
//	condition["f_folder_id"] = "13"
//	condition["f_name"] = "filex.txt"
//	data, err := dao.GetOne("file", 8, condition)
//	checkError(err)
//	log.Println(data)
//}

//func TestGetListColumn(t *testing.T) {
//	columns := []string{"f_name", "f_stuffix", "f_size"}

//	condition := make(map[string]interface{})
//	condition["f_id"] = "16"
//	condition["f_folder_id"] = "13"
//	condition["f_name"] = "filex.txt"

//	data, err := dao.GetListColumn("file", columns, condition)
//	checkError(err)
//	for _, v := range data {
//		log.Println(v)
//	}
//}

//func TestGetOneColumn(t *testing.T) {
//	columns := []string{"f_name", "f_stuffix", "f_size"}

//	condition := make(map[string]interface{})
//	condition["f_id"] = "16"
//	condition["f_folder_id"] = "13"
//	condition["f_name"] = "filex.txt"

//	data, err := dao.GetListColumn("file", columns, condition)
//	checkError(err)
//	log.Println(data)
//}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
