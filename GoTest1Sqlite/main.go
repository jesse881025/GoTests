package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"os"
	"path"
	"GoTest1Sqlite/models"
)


//常量的定义
const (
	_DB_NAME = "data/goTest1Sqlite.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func init() {
	fmt.Println("main.go - init  start")

	//判断文件是否存在，不存在则先创建
	if !IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	// 参数1   driverName
	// 参数2   数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME,10)

	orm.Debug = true
	orm.RunSyncdb("default", false, true)


	fmt.Println("main.go - init  end")
}

//判断文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	//或者
	//return err == nil || !os.IsNotExist(err)
	//或者
	//return !os.IsNotExist(err)
}


func main() {
	fmt.Println("main.go - main  start")

	//添加数据
	//addData()
	//查询数据
	//queryData()
	//删除数据
	deletData()

	fmt.Println("main.go - main  end")
}


//添加数据
func addData()  {

	models.AddUser("zhang1")

	user := &models.User{Name:"zhang2"}
	models.AddUserObject(user)

	user3 := &models.User{Name:"zhang3"}
	user3.Profile = &models.Profile{Age:10}
	models.AddUserObject(user3)

}

//查询数据
func queryData()  {
	user,err := models.ReadUser(3)
	if err == nil {
		fmt.Println(user)
	}

	user1,err := models.ReadUser2(3)
	if err == nil {
		fmt.Println(user1)
	}

	profile,err := models.ReadUserProfile(3)
	if err == nil {
		fmt.Println(profile)
	}
}

//删除数据
func deletData()  {
	err := models.DeletUser(3)
	if err != nil {
		fmt.Println(err)
	}
}











