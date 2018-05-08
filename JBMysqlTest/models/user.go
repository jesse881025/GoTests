package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id 		int
	Name	string	`orm:"size(100)"`
	//Age 	int
	//Nick    string
}

func init()  {

	// 注册模型
	orm.RegisterModel(new(User))

}

func AddUser(name string) error {
	// 获取orm对象
	o := orm.NewOrm()
	// 创建一个 User 对象
	user := &User{Name: name}

	//查询name是否被用了
	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user)
	if err == nil {
		// 查询到了，如果有则不在插入
		return err
	}

	// 如果不存在则插入
	//var id int
	_, err = o.Insert(user)

	return err
}

func UpdateUser(id int, name string) error {
	// 获取orm对象
	o := orm.NewOrm()
	// 创建一个 User 对象
	user := &User{Id: id}
	user.Name = name

	// 更新
	_, err := o.Update(user)

	return err
}

func DeletUser(id int) error  {
	// 获取orm对象
	o := orm.NewOrm()
	// 创建一个 User 对象
	user := &User{Id: id}

	// 删除
	_, err := o.Delete(user)

	return err
}

func ReadUser(id int) (User, error) {
	// 获取orm对象
	o := orm.NewOrm()
	// 创建一个 User 对象
	user := &User{Id: id}

	// 读取
	err := o.Read(user)

	return *user, err
}




















