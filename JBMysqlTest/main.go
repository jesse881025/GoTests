package main

import (
	"github.com/astaxie/beego/orm"
	"JBMysqlTest/models"
	"time"
	"github.com/astaxie/beego"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

const (
	//_DB_NAME = "data/jbMysqlTest"
	//_DB_NAME = "jbMysqlTest"
	_DB_NAME = "jbmysqltest"
)

func init()  {

	// 注册模型
	//orm.RegisterModel(new(models.User))

	// 注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)//可省略，默认已注册

	// 注册默认数据库
	// 我的mysql的root用户密码为123456，打算把数据表建立在名为 jbMysqlTest 数据库里
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
	str := "root:123456@tcp(127.0.0.1:3306)/" + _DB_NAME + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", str,30)


	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)

}



func main() {

	var err error

	err = models.AddUser("红伟")
	if err != nil {
		beego.Error(err)
	}

	user, err := models.ReadUser(1)
	if err != nil {
		beego.Error(err)
	}
	fmt.Println(user)

	err = models.UpdateUser(user.Id, "红伟222")
	if err != nil {
		beego.Error(err)
	}


	fmt.Println(user)
	user, err = models.ReadUser(1)
	if err != nil {
		beego.Error(err)
	}
	fmt.Println(user)

	err = models.DeletUser(user.Id)
	if err != nil {
		beego.Error(err)
	}

}




















