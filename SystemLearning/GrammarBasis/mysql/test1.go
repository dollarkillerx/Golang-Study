package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

//创建数据库中的映射
type Person struct {
	Name string `db:"name"`
	Age int `db:"age"`
	Money float64 `db:"rmb"`
}

func main() {
	//curl1()
	//curd2()
	curd3()
}


func curl1()  {
	//设置时区
	l,_ := time.LoadLocation("Asia/Shanghai")

	db, err := sqlx.Open("mysql", "beego:beego@tcp(127.0.0.1:3306)/beego")
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		db.Close()
	}()

	time.Now().In(l).Unix()
	//curl
	//result,err := db.Exec("INSERT INTO persion(name,age,rmb,gender,birthday) VALUE(?,?,?,?,?);","张全蛋",30,100,1,time.Now().In(l).Unix())
	//result, err := db.Exec("DELETE FROM persion WHERE name like ?;", "%蛋")
	result, err := db.Exec("UPDATE `persion` SET `name` = ? WHERE `name` like ?;", "大公鸡", "%蛋")
	if err!=nil {
		panic(err.Error())
	}
	rowAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	fmt.Println("多少行收到了影响",rowAffected)
	fmt.Println("最后插入一行的id",lastInsertId)
}

func curd2() {
	//设置时区域
	location, _ := time.LoadLocation("Asia/Shanghai")
	db, e := sqlx.Open("mysql", "beego:beego@tcp(127.0.0.1:3306)/beego")
	if e!=nil {
		panic(e.Error())
	}
	defer func() {db.Close()}()
	time.Now().In(location).Unix()

	//预定义person切片用户接收查询结果
	ps := make([]Person,0)

	//执行查询
	e = db.Select(&ps, "SELECT name,age,rmb FROM persion")
	if e != nil {
		panic(e.Error())
	}
	for _,k := range ps {
		//arr := [...]int{1,2,3,4,5}
		fmt.Println(k.Age)
	}

	fmt.Println("查询成功",ps)

}

func curd3() {
	db, e := sqlx.Connect("mysql", "beego:beego@tcp(127.0.0.1:3306)/beego")
	if e!=nil {
		panic(e.Error())
	}
	defer func() {
		db.Close()
	}()

	//开启事物
	tx, _ := db.Begin()

	//执行系列的增删改
	result, e := tx.Exec("INSERT INTO persion(`name`,`age`,`gender`) VALUE (?,?,?)", "吕亚丹", 20, 0)

	//判断是否出错,出错则回滚
	if e!=nil{
		//回滚
		tx.Rollback()
		panic(e.Error())
	}else{
		//提交
		tx.Commit()
	}
	i, _ := result.RowsAffected()
	fmt.Println("影响行数",i)

}

/**
CREATE TABLE persion(
    `id` int AUTO_INCREMENT,
    `name` varchar(20) not null DEFAULT '',
    `age` int,
    `rmb` float,
    `gender` tinyint,
    `birthday` int,
    PRIMARY KEY(`id`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8;
 */