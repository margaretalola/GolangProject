package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	kd_mk  string
	matkul string
	hari   string
	masuk  string
	keluar string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mahasiswa")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//var age = 27
	rows, err := db.Query("select * from tbl_mhs")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.kd_mk, &each.matkul,
			&each.hari, &each.masuk, &each.keluar)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each)
		//fmt.Println(each.matkul)
		//fmt.Println(each.id)
	}
}

func main() {
	sqlQuery()
}
