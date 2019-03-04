package models

import
(
    "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"time"
)

/*
 * user表初始化
 *@auth heyibo
 *@date 2018-5-30
 */
type User struct {
	Id                   int      `orm:"pk:auto"`
	Username             string   `orm:"size(255)"`
	Password             string   `orm:"size(255)"`
	Name                 string   `orm:"size(255)"`    //用户名
	BirthDate            string   `orm:"size(255)"`    //生日
	Perjectid            int      `orm:"size(10)"`
	Status               int      `orm:"size(10)"`
	Gender               int      `orm:"size(8)"`      //性别
	Email                string   `orm:"size(255)"`    //Email
	Phone                string   `orm:"size(255)"`    //电话
	CreateTime              time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdateTime              time.Time   `orm:"auto_now;type(datetime)"`
}

/*
 * 查询所有用户方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *User) GetAll() (dataList []*User,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*User
	qs:= o.QueryTable("user")
	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}


/*
 * 添加用户方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *User) AddOne(User *User) (id int,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Insert(User); err == nil {
		logs.Info(" Insert in database:", result)
		return User.Id, nil
	}
	return 0, err
}




/*
 * 修改用户方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c * User) Update(User *User) (err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Update(User); err == nil {
		logs.Info(" updated in database:", result)
	}
	return
}

/*
 * 删除用户方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *User) Delete(id int) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=User{Id:id}
	if err = o.Read(&u, "Id"); err == nil {
		if num, err := o.Delete(&User{Id: id}); err == nil {
			logs.Info(" deleted in usertable:", num)
		}
	}
	return
}


/*
 * 根据ID查询方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *User) GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=&User{Id: id}
	if err = o.Read(u, "Id"); err == nil {
		return u, nil
	}
	return nil, err
}
/*
 * 根据ID查询方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *User) GetUserByName(Name string) (dataList []*User,err error ) {
	o := orm.NewOrm()
	o.Using("default")
	var list []*User
	name :=Name
	qs := o.QueryTable("user").Filter("name__icontains",name)
	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}
/*
 * 根据用户名查询密码方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *User) LoginUser(username string, password string) ( user []*User,err error,) {
	o := orm.NewOrm()
	o.Using("default")
	var users []*User
	qs := o.QueryTable("user").Filter("username",username).Filter("password",password).Filter("status",1).GroupBy("Id")
	//查询数据
	if _, err = qs.All(&users,"Id"); err == nil {
		for _, v := range users {
			user = append(user, v)
		}
		return user, nil
	}
	logs.Info("user :",user)
	return user, err
}

/*
 * 数据库初始化方法
 *@auth heyibo
 *@date 2018-5-30
 */
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
	//创建表
	//orm.RunSyncdb("default", false, true)
}