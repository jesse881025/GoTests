package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id 			int64
	Name 		string
	Profile		*Profile	`orm:"rel(one)"`  //设置一对一关系
}

func init()  {
	fmt.Println("user.go - init  start")


	//将你定义的 Model 进行注册，最佳设计是有单独的 models.go 文件，
	//在他的 init 函数中进行注册。
	orm.RegisterModel(new(User))

	fmt.Println("user.go - init  end")
}

func AddUser(name string) (int64, error) {
	//创建一个User对象
	user := &User{Name:name}
	id, err := AddUserObject(user)
	return id, err
}

func AddUserObject(user *User) (int64, error) {
	//获取orm对象
	o := orm.NewOrm()
	//创建一个User对象
	userTemp := &User{Name:user.Name}

	var id int64 = 0
	//查询name是否被用了
	qs := o.QueryTable("user")
	err := qs.Filter("name",userTemp.Name).One(userTemp)
	if err == nil {
		return id, err
	}

	//检查 profile, 不能为空
	err = checkUserProfile(user)
	if err != nil {
		return id, err
	}

	//如果不存在则插入
	id, err = o.Insert(user)
	if err != nil {
		return id, err
	}
	return id, nil
}


//检查 profile, 不能为空
func checkUserProfile(user *User) error {
	if user.Profile == nil {
		user.Profile = &Profile{Age:0}
	}
	id, err := AddProfileObject(user.Profile)
	if err == nil {
		user.Profile.Id = id
	}
	return err
}



//根据 ID 查询, 关系查询
//方法一：
func ReadUser(id int64) (*User, error) {
	//获取orm对象
	o := orm.NewOrm()
	//创建一个User对象
	userTemp := &User{Id:id}
	err := o.Read(userTemp)
	if userTemp.Profile != nil {
		err = o.Read(userTemp.Profile)
	}
	return userTemp, err
}
//方法二： 直接关联查询：
func ReadUser2(id int64) (*User, error) {
	//获取orm对象
	o := orm.NewOrm()
	//创建一个User对象
	userTemp := &User{Id:id}
	//自动查询到 Profile
	//因为在 Profile 里定义了反向关系的 User，所以 Profile 里的 User 也是自动赋值过的，可以直接取用。
	err := o.QueryTable("user").Filter("Id", id).RelatedSel().One(userTemp)
	return userTemp, err
}

//通过 User 反向查询 Profile
func ReadUserProfile(id int64) (Profile, error) {
	var profileTemp Profile
	//获取orm对象
	o := orm.NewOrm()
	//之前创建 Profile 表时，加了 prefix_ 的前缀
	err := o.QueryTable("prefix_profile").Filter("User__Id", id).One(&profileTemp)
	return profileTemp, err
}

func DeletUser(id int64) error {
	//获取orm对象
	o := orm.NewOrm()
	//创建一个User对象
	userTemp := &User{Id:id}
	_, err := o.Delete(userTemp)
	return err
}


