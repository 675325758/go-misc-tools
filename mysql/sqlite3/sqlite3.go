package sqlite3

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func test() {
	fmt.Println("enter sqlite3 test")
	//open
	db, err := sql.Open("sqlite3", "/tmp/my.db")
	if err != nil {
		panic(err)
	}
	defer func() {
		db.Close()
	}()

	//create table
	table := `
    CREATE TABLE IF NOT EXISTS user (
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(128) NULL,
        created DATE NULL
    );
    `
	_, err = db.Exec(table)
	if err != nil {
		panic(err)
	}
	//insert
	stmt, err := db.Prepare("INSERT INTO user(name,  created) values(?,?)")
	if err != nil {
		panic(err)
	}
	// res 为返回结果
	res, err := stmt.Exec("guoke", "2012-12-09")
	if err != nil {
		panic(err)
	}

	// 可以通过res取自动生成的id
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("id=", id)

	//delete
	stmt, err = db.Prepare("delete from user where uid=?")
	if err != nil {
		panic(err)
	}

	res, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	//up
	stmt, err = db.Prepare("update user set name=? where uid=?")
	if stmt == nil || err != nil {
		panic(err)
	}
	_, err = stmt.Exec("guoke3915", 1)
	if err != nil {
		panic(err)
	}
	//insert2
	stmt, err = db.Prepare("INSERT INTO user(name,  created) values(?,?)")
	if err != nil {
		panic(err)
	}
	// res 为返回结果
	res, err = stmt.Exec("guoke", "2012-12-09")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var uid int
		var name string
		var created time.Time
		err = rows.Scan(&uid, &name, &created)
		if err != nil {
			panic(err)
		}

		fmt.Println(uid, name, created)
	}
}

func Test() {
	go test()
}
