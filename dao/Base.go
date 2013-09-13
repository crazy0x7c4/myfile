package dao

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	mysqlDB, _ := sql.Open("mysql", "crazy:12345@/myfile?charset=utf8")
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

func Add(table string, values ...interface{}) (int, error) {
	str := bytes.Buffer{}
	str.WriteString("INSERT INTO ")
	str.WriteString(table)
	str.WriteString(" VALUES (")

	l := len(values)
	for i := 0; i < l; i++ {
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(",")
		} else {
			str.WriteString(")")
		}
	}

	id, err := Execute(str.String(), values...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Delete(table string, condition map[string]interface{}) error {
	str := bytes.Buffer{}
	str.WriteString("DELETE FROM ")
	str.WriteString(table)
	str.WriteString(" WHERE ")

	var values []interface{}
	l := len(condition)
	i := 0
	for k, v := range condition {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(" AND ")
		}
		i = i + 1
	}

	_, err := Execute(str.String(), values...)
	if err != nil {
		return err
	}

	return nil
}

func Update(table string, columns map[string]interface{}, condition map[string]interface{}) error {
	str := bytes.Buffer{}
	str.WriteString("UPDATE ")
	str.WriteString(table)
	str.WriteString(" SET ")

	var values []interface{}
	l := len(columns)
	i := 0
	for k, v := range columns {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(", ")
		}
		i = i + 1
	}

	str.WriteString(" WHERE ")

	l = len(condition)
	i = 0
	for k, v := range condition {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(" AND ")
		}
		i = i + 1
	}

	_, err := Execute(str.String(), values...)
	if err != nil {
		return err
	}

	return nil
}

func GetList(table string, condition map[string]interface{}) ([][]string, error) {
	str := bytes.Buffer{}
	str.WriteString("SELECT * FROM ")
	str.WriteString(table)

	var values []interface{}
	l := len(condition)
	if l > 0 {
		str.WriteString(" WHERE ")
	}
	i := 0
	for k, v := range condition {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(" AND ")
		}
		i = i + 1
	}

	data, err := Query(str.String(), values...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetOne(table string, columns int, condition map[string]interface{}) ([]string, error) {
	str := bytes.Buffer{}
	str.WriteString("SELECT * FROM ")
	str.WriteString(table)

	var values []interface{}
	l := len(condition)
	if l > 0 {
		str.WriteString(" WHERE ")
	}
	i := 0
	for k, v := range condition {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(" AND ")
		}
		i = i + 1
	}

	data, err := QueryRow(columns, str.String(), values...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetListColumn(table string, columns []string, condition map[string]interface{}) ([][]string, error) {
	str := bytes.Buffer{}
	str.WriteString("SELECT ")

	l := len(columns)

	for i := 0; i < l; i++ {
		str.WriteString(columns[i])
		if i < l-1 {
			str.WriteString(", ")
		}
	}

	str.WriteString(" FROM ")
	str.WriteString(table)

	var values []interface{}
	l = len(condition)
	if l > 0 {
		str.WriteString(" WHERE ")
	}
	i := 0
	for k, v := range condition {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < l-1 {
			str.WriteString(" AND ")
		}
		i = i + 1
	}

	data, err := Query(str.String(), values...)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetOneColumn(table string, columns []string, condition map[string]interface{}) ([]string, error) {
	str := bytes.Buffer{}
	str.WriteString("SELECT ")

	colLen := len(columns)

	for i := 0; i < colLen; i++ {
		str.WriteString(columns[i])
		if i < colLen-1 {
			str.WriteString(", ")
		}
	}

	str.WriteString(" FROM ")
	str.WriteString(table)

	var values []interface{}
	conLen := len(condition)
	if conLen > 0 {
		str.WriteString(" WHERE ")
	}
	i := 0
	for k, v := range condition {
		values = append(values, v)
		str.WriteString(k)
		str.WriteString("=")
		str.WriteString("?")
		if i < conLen-1 {
			str.WriteString(" AND ")
		}
		i = i + 1
	}

	data, err := QueryRow(colLen, str.String(), values...)
	if err != nil {
		return nil, err
	}
	return data, nil
}
