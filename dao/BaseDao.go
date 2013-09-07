package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	mysqlDB, _ := sql.Open("mysql", "root:12345@/test?charset=utf8")
	db = mysqlDB
}

func Execute(sqlStr string, params ...interface{}) (int, error) {
	result, err := db.Exec(sqlStr, params...)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func Query(sqlStr string, params ...interface{}) ([][]string, error) {
	rows, err := db.Query(sqlStr, params...)
	if err != nil {
		return nil, err
	}

	cols, _ := rows.Columns()

	colsLen := len(cols)

	datas := [][]string{}

	for rows.Next() {
		data := make([]sql.RawBytes, colsLen)
		scanArgs := make([]interface{}, colsLen)

		for i := 0; i < colsLen; i++ {
			scanArgs[i] = &data[i]
		}
		rows.Scan(scanArgs...)

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

	return datas, nil
}

func QueryRow(columnNum int, sqlStr string, params ...interface{}) ([]string, error) {
	row := db.QueryRow(sqlStr, params...)

	data := make([]string, columnNum)
	scanArgs := make([]interface{}, columnNum)
	for i := 0; i < columnNum; i++ {
		scanArgs[i] = &data[i]
	}
	err := row.Scan(scanArgs...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
