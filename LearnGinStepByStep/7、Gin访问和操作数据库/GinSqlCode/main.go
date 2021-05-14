package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	//1、链接数据库
	connStr := "root:12345678@tcp(127.0.0.1:3306)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建数据库表
	//person: id,name,age
	//_, err = db.Exec("create table person(" +
	//	"id int auto_increment primary key," +
	//	"name  varchar(12) not null," +
	//	"age int default 1" +
	//	");")
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//} else {
	//	fmt.Println("数据库表创建成功")
	//}

	//插入数据到数据库表
	_, err = db.Exec("insert into person(name,age) "+
		"values(?,?);", "Jack", 20)

	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据插入成功")
	}

	//查询数据库
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

scan:
	if rows.Next() {
		person := new(Person)
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}

	//->
	//Davie  18
	//Tom   15
	//Jack  20
	//Lily 25
	//->
}

type Person struct {
	Id   int
	Name string
	Age  int
}
