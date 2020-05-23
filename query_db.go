package dbtoapi

import (
	"database/sql"
	"fmt"
	_ "github.com/goees/go-mysql"
	"strconv"
)

//查询一共有哪些表
func QueryDBTables() []map[string]string {
	db, err := openDB()
	checkErr(err)
	if err == nil {
		//没有报错，则关闭DB释放连接
		defer closeDB(db)
	}

	var data []map[string]string
	if err == nil {
		data = queryTables(db)
	}
	return data
}

//查询指定表并返回，分页
func QueryDB(table string, pageNum int, pageSize int) []map[string]string {
	db, err := openDB()
	checkErr(err)
	if err == nil {
		//没有报错，则关闭DB释放连接
		defer closeDB(db)
	}

	var data []map[string]string
	if err == nil {
		data = query(db, table, pageNum, pageSize)
	}
	return data
}

//连接MySQL/GBase
func openDB() (*sql.DB, error) {
	//格式：
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//从配置文件读取
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", conf.DBUsername, conf.DBPassword, conf.DBServerIP, conf.DBServerPort, conf.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	return db, err
}

//关闭DB
func closeDB(db *sql.DB) {
	closeErr := db.Close()
	checkErr(closeErr)
}

//查询一共有哪些表
func queryTables(db *sql.DB) []map[string]string {
	return getData(db, "show tables")
}

//查询指定表
func query(db *sql.DB, table string, pageNum int, pageSize int) []map[string]string {
	//拼接SQL
	a, b := (pageNum-1)*pageSize, pageSize
	querySql := "SELECT * FROM " + table
	if a == 0 {
		querySql += " LIMIT " + strconv.Itoa(b)
	} else {
		querySql += " LIMIT " + strconv.Itoa(a) + "," + strconv.Itoa(b)
	}
	return getData(db, querySql)
}

//查询表，返回结果集(切片类型)
func getData(db *sql.DB, querySql string) []map[string]string {
	//查询
	rows, err := db.Query(querySql)
	checkErr(err)

	//读出查询出的列字段名
	cols, _ := rows.Columns()

	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))

	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))

	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	//最后得到的map
	//results := make(map[int]map[string]string)
	//改成切片（保存查询结果的切片）
	results := make([]map[string]string, 0)

	i := 0
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			checkErr(err)
			return nil
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		//装入结果集中
		//results[i] = row
		results = append(results, row)
		i++
	}

	return results
}
