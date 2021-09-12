package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"os"
)


func queryPlayer(playerName string) (interface{}, error) {
	db, err := sqlx.Open("mysql","root:admin@tcp(127.0.0.1:3306)/weiqi?charset=utf8")
	if err != nil {
		return nil, errors.Wrap(err, "mysql connect failed")
	}

	type userInfo struct {
		Id int `db:"id"`
		Name string `db:"name"`
		Birthday string `db:"birthday"`
		TeamId int `db:"teamId"`
	}

	//初始化定义结构体，用来存放查询数据
	var userData *userInfo = new(userInfo)
	err = db.Get(userData,fmt.Sprintf("select * from player where name = '%s'", playerName))
	if err != nil {
		return nil, errors.Wrap(err, "query op error")
	}

	return userData, nil
}

func main() {
	// dao 层中当遇到一个 sql.ErrNoRows 的时候，应该 Wrap 这个 error，抛给上层
	// 因为上层调用方需要知道NoRows的具体信息、是sql语句错了、还是正常的没有数据，以便进行错误处理
	userData, err := queryPlayer("KeJie")
	if err != nil {
		fmt.Printf("orginal error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}

	fmt.Println("userData:", userData)
}

