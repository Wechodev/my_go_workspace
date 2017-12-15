package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/aliyoyo_com?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT cf_api_ip SET api_id=?, ip=?")

	res, err := stmt.Exec("20", "192.168.0.0")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	stmt, err = db.Prepare("UPDATE cf_api_ip SET api_id=?, ip=? WHERE api_ip_id=?")
	checkErr(err)

	res, err = stmt.Exec("911", "128.0.0.1",  id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM cf_api_ip")
	checkErr(err)

	for rows.Next() {
		var api_ip_id int
		var api_id int
		var ip string
		err = rows.Scan(&api_ip_id, &api_id, &ip)
		checkErr(err)
		fmt.Println(api_ip_id)
		fmt.Println(api_id)
		fmt.Println(ip)
	}

	stmt, err = db.Prepare("DELETE  FROM cf_api_ip WHERE api_ip_id=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}
