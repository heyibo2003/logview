package models

import
(
"github.com/astaxie/beego/orm"
"github.com/astaxie/beego/logs"
	"time"
)

/*
 * 初始化项目表
 *@auth heyibo
 *@date 2018-5-30
 */
type Perject struct {
	Id                     int      `orm:"pk:auto"`
	Perject                string   `orm:"size(255)"`
	CreateTime              time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdateTime              time.Time   `orm:"auto_now;type(datetime)"`
}


func (this *Perject) TableName() string {
	return "perject"
}

/*
 * 查询所有perject方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *Perject) GetAll() (dataList []*Perject,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Perject
	qs:= o.QueryTable("perject")
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
 * 添加项目名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Perject) Add(Perject *Perject) (id int,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Insert(Perject); err == nil {
		logs.Info(" Insert in database:", result)
		return Perject.Id, nil
	}
	return Perject.Id, err
}

/*
 * 修改项目名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c * Perject) Update(Perject *Perject) (err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	o.Update(Perject)
	return
}



/*
 * 删除项目名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Perject) Delete(id int) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	u:=Perject{Id:id}
	if err = o.Read(&u, "Id"); err == nil {
		if num, err := o.Delete(&Perject{Id: id}); err == nil {
			logs.Info(" deleted in perjecttable:", num)
		}
	}
	return
}

/*
 * 根据ID查询方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Perject) GetPerjectById(id int) (v *Perject, err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=&Perject{Id: id}
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
	orm.RegisterModel(new(Perject))
	//创建表
	//orm.RunSyncdb("default", false, true)
}