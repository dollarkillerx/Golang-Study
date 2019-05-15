package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//创建数据库中的映射
type Person struct {
	Name string `db:"name"`
	Age int `db:"age"`
	Money float64 `db:"rmb"`
}

func main() {
	sqlx.Open("mysql","root:")
}