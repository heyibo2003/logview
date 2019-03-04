package models

import
(
"github.com/astaxie/beego/orm"
"github.com/astaxie/beego/logs"
	"time"
)

/*
 * 初始化环境表
 *@auth heyibo
 *@date 2018-5-30
 */
type Env struct {
	Id        int      `orm:"pk:auto"`
	Envname   string   `orm:"size(255)"`
	CreateTime      time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time   `orm:"auto_now;type(datetime)"`
}


func (this *Env) TableName() string {
	return "env"
}
/*
 * 查询所有用户方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *Env) GetAll() (dataList []*Env,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Env
	qs:= o.QueryTable("env")
	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}

/*
 * 添加环境名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Env) Add(Env *Env) (id int,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Insert(Env); err == nil {
		logs.Info(" Insert in database:", result)
		return Env.Id, nil
	}
	return Env.Id, err
}


/*
 * 修改环境名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c * Env) Update(Env *Env) (err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	o.Update(Env)
	return
}


/*
 * 删除环境名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c * Env)  Delete(id int) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=Env{Id:id}
	if err = o.Read(&u, "Id"); err == nil {
		if num, err := o.Delete(&Env{Id: id}); err == nil {
			logs.Info(" deleted in envttable:", num)
		}
	}
	return
}
/*
 * 根据ID查询方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Env) GetEnvById(id int) (v *Env, err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=&Env{Id: id}
	if err = o.Read(u, "Id"); err == nil {
		return u, nil
	}
	return nil, err
}

/*
 * 数据库初始化方法
 *@auth heyibo
 *@date 2018-5-30
 */
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Env))
	//创建表
	//orm.RunSyncdb("default", false, true)
}