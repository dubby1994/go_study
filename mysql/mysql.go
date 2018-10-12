package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
	"fmt"
)

func (data Data) String() string {
	return fmt.Sprintf("id:%d key:%s value:%s\n[createTime:%d updateTime:%d]\n",
		data.Id, data.Key, data.Value, data.CreateTime.Unix(), data.UpdateTime.Unix())
}

type Data struct {
	Id         int64
	Key        string
	Value      string
	CreateTime time.Time //[]uint8
	UpdateTime time.Time //[]uint8
}

func main() {
	db, err := sql.Open("mysql", "dubby:123456@tcp(127.0.0.1:3306)/go_test?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("========QueryById========")
	data := QueryById(2, db)
	fmt.Println(*data)

	fmt.Println("========QueryByIdRange========")
	resultList := QueryByIdRange(1, 10, db)
	for i := 0; i < len(resultList); i++ {
		fmt.Println(*resultList[i])
	}

	fmt.Println("========Insert========")
	data = Insert("go_key", "go_value", db)
	fmt.Println(*data)

	fmt.Println("========UpdateById========")
	data = QueryById(2, db)
	prefix := fmt.Sprintf("update-%d-", time.Now().UnixNano()/1000%10000)
	data = UpdateById(data.Id, prefix+"dubby", prefix+"www.dubby.cn", db)
	fmt.Println(*data)
}

func Insert(key string, value string, db *sql.DB) *Data {
	stmtOut, err := db.Prepare("INSERT INTO `data` (`key`, `value`) values (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	result, err := stmtOut.Exec(key, value)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return QueryById(id, db)
}

func UpdateById(id int64, key string, value string, db *sql.DB) *Data {
	stmtOut, err := db.Prepare("UPDATE `data` SET `key`=?, `value`=? WHERE `id`=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	_, err = stmtOut.Exec(key, value, id)
	if err != nil {
		panic(err.Error())
	}

	return QueryById(id, db)
}

func QueryById(idRequest int64, db *sql.DB) *Data {
	stmtOut, err := db.Prepare("SELECT * FROM `data` WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows := stmtOut.QueryRow(idRequest)

	data := new(Data)
	err = rows.Scan(&data.Id, &data.Key, &data.Value, &data.CreateTime, &data.UpdateTime)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func QueryByIdRange(minId int64, maxId int64, db *sql.DB) []*Data {
	stmtOut, err := db.Prepare("SELECT * FROM `data` WHERE id >= ? AND id <= ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(minId, maxId)
	if err != nil {
		panic(err.Error())
	}

	var result []*Data

	for rows.Next() {
		data := new(Data)
		err = rows.Scan(&data.Id, &data.Key, &data.Value, &data.CreateTime, &data.UpdateTime)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, data)
	}

	return result
}
