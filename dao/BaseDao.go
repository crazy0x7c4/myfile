package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	mysqlDB, err := sql.Open("mysql", "root:12345@/test?charset=utf8")
	checkError(err)
	db = mysqlDB
}

func Execute(sqlStr string, params ...interface{}) uint {
	result, err := db.Exec(sqlStr, params...)
	checkError(err)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Execute SQL Error:", err)
		}
	}()

	id, err := result.LastInsertId()
	checkError(err)

	return uint(id)
}

func Query(sqlStr string, params ...interface{}) [][]string {
	rows, err := db.Query(sqlStr, params...)
	checkError(err)
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Query Error:", err)
		}
	}()

	cols, err := rows.Columns()
	checkError(err)

	colsLen := len(cols)

	datas := [][]string{}

	for rows.Next() {
		data := make([]sql.RawBytes, colsLen)
		scanArgs := make([]interface{}, colsLen)

		for i := 0; i < colsLen; i++ {
			scanArgs[i] = &data[i]
		}
		err := rows.Scan(scanArgs...)
		checkError(err)

		dataStr := make([]string, colsLen)
		for k, v := range data {
			if v == nil {
				dataStr[k] = ""
			} else {
				dataStr[k] = string(v)
			}
		}

		datas = append(datas, dataStr)
	}
	defer rows.Close()

	return datas
}

func QueryRow(columnNum int, sqlStr string, params ...interface{}) []string {
	row := db.QueryRow(sqlStr, params...)

	data := make([]string, columnNum)
	scanArgs := make([]interface{}, columnNum)
	for i := 0; i < columnNum; i++ {
		scanArgs[i] = &data[i]
	}
	err := row.Scan(scanArgs...)
	checkError(err)

	return data
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
