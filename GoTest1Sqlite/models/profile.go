package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id 		int64
	Age 	int16
	User 	*User	`orm:"reverse(one)"`  // 设置一对一反向关系(可选)
}

func init()  {
	fmt.Println("profile.go - init  start")

	//将你定义的 Model 进行注册，最佳设计是有单独的 models.go 文件，
	//在他的 init 函数中进行注册。
	//orm.RegisterModel(new(Profile))

	//使用表名前缀,创建后的表名为 prefix_profile
	orm.RegisterModelWithPrefix("prefix_", new(Profile))


	fmt.Println("profile.go - init  end")
}

func AddProfile(age int16) (int64, error)  {
	o := orm.NewOrm()
	profile := &Profile{Age:age}
	id, err := o.Insert(profile)
	return id, err
}

func AddProfileObject(profile *Profile) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(profile)
	return id, err
}

func DeletProfile(id int64) error {
	//获取orm对象
	o := orm.NewOrm()
	//创建一个User对象
	profile := &Profile{Id:id}
	_, err := o.Delete(profile)
	return err
}