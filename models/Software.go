package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"fmt"
)

type Software struct {
	Id        int    `orm:"pk:auto"`
	Domain    string `orm:"size(255)"`
	Checkscript string `orm:"size(255)"`
	Script    string `orm:"size(255)"`
	Path      string `orm:"size(255)"`
	Port      int    `orm:"size(10)"`
	Softname  string `orm:"size(100)"`
	Softpath  string `orm:"size(100)"`
	EnvId     int    `orm:"size(10)"`
	PerjectId int    `orm:"size(10)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}
type softlist struct {
	Id        int
	Domain    string
	Script    string
	Checkscript string
	Path      string
	Port      int
	Softname  string
	Softpath  string
	Envname   string
	Perject   string
	CreateTime time.Time
	UpdateTime time.Time
}

type Execlist struct {
	Id        int
	Domain    string
	Script    string
	Checkscript string
	Path      string
	Port      int
	Softname  string
	Softpath  string
	Machine   string
	User      string
	Passwd    string
	Hostport  int
}

type  perlist struct {
	Id        int
	Perject   string
}

type Envlist struct {
	Id    int
	Envname   string
}


func (this *Software) TableName() string {
	return "software"
}



/*
 * 查询所有用户方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *Software) GetAll() (dataList []*softlist,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*softlist
	qs,err:= o.Raw("SELECT software.`id`,software.`domain`,software.`script`,software.`checkscript`,software.`path`,software.`port`,software.`softname`,software.`softpath`,env.`envname`,perject.`perject`,software.`Create_Time`,software.`Update_Time` FROM software INNER JOIN env,perject WHERE software.`Env_Id` =env.`id` AND software.`Perject_Id` = perject.`id`").QueryRows(&list)
	//查询数据
	//查询数据software.`port`
	if  err == nil &&qs>0{
		for _, v := range list {
			//fmt.Println("123",v.UpdateTime)
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err

}


/*
 * 关联查询，查询项目ID,和项目名称，根据项目ID去重复
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Software) GetPerjectList() (dataList []*perlist,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*perlist
	qs,err:= o.Raw("SELECT perject.`id`,perject.`perject` FROM software INNER JOIN perject WHERE  software.`Perject_Id` = perject.`id`  GROUP BY software.`Perject_Id`").QueryRows(&list)
	//查询数据
	//查询数据software.`port`
	if  err == nil &&qs>0{
		for _, v := range list {
			fmt.Println("qs:",v.Id,v.Perject)
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}
/*
* 关联查询，根据项目ID,查询环境ID 和环境名称,根据环境ID去重复
*@auth heyibo
*@date 2018-5-30
*/
func (c *Software) GetEnvList(PerjectId int) (dataList []*Envlist,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Envlist
	qs,err:= o.Raw("SELECT env.`id`,env.`envname` FROM software INNER JOIN env WHERE  software.`Env_Id` = env.`id` AND software.`Perject_Id` = ? GROUP BY software.`Env_Id`",PerjectId).QueryRows(&list)
	//查询数据
	//查询数据software.`port`
	if  err == nil &&qs>0{
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}

/*
 * 添加方法 ,根据perjectId,EnvId查询domain的方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Software) GetSoftwareByPerjectIdAndEnvId (PerjectId int,EnvId int) (dataList []*Software,err error )  {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Software
	qs := o.QueryTable("software").Filter("PerjectId",PerjectId).Filter("EnvId",EnvId).GroupBy("domain")
	//查询数据
	if _, err = qs.All(&list,"domain"); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-domain :", dataList)
		return dataList, nil

	}
	return dataList, err
}



/*
 * 添加方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *Software) AddOne(Software *Software) (id int,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Insert(Software); err == nil {
		logs.Info(" Insert in database:", result)
		return Software.Id, nil
	}
	return 0, err
}




/*
 * 修改方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c * Software) Update(Software *Software) (err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Update(Software); err == nil {
		logs.Info(" updated in database:", result)
	}
	return
}

/*
 * 删除方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Software) Delete(id int) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=Software{Id:id}
	if err = o.Read(&u, "Id"); err == nil {
		if num, err := o.Delete(&Software{Id: id}); err == nil {
			logs.Info(" deleted in softwaretable:", num)
		}
	}
	return
}


/*
 * 根据ID查询方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Software) GetSoftwareById(id int) (v *Software, err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=&Software{Id: id}
	if err = o.Read(u, "Id"); err == nil {
		return u, nil
	}
	return nil, err
}


/*
 * 根据domain查询方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Software) GetSoftwareByDomain(Domain string) (dataList []*softlist,err error ) {
	o := orm.NewOrm()
	o.Using("default")
	var list []*softlist
	domain :=Domain
	qs,err:=o.Raw("SELECT software.`id`,software.`domain`,software.`script`,software.`path`,software.`port`,software.`softname`,software.`softpath`,env.`envname`,perject.`perject`,software.`Create_Time`,software.`Update_Time` FROM software INNER JOIN env,perject WHERE software.`Env_Id` =env.`id` AND software.`Perject_Id` = perject.`id` AND software.`domain` LIKE '%"+domain+"%'").QueryRows(&list)
	//查询数据
	if  err == nil && qs>0{
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}
/*
 * 添加方法根据perjectId,EnvId,domain,machine等字段关联查询software和Environmentals表
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Software) GetSoftwareByEnvironmentals(PerjectId int,EnvId int,Domain string,Machine string) (dataList []*Execlist,err error ) {
	o := orm.NewOrm()
	o.Using("default")
	var list []*Execlist
	qs,err:=o.Raw("SELECT DISTINCT software.`id`,software.`domain`,software.`script`,software.`checkscript`,software.`path`,software.`port`,software.`softname`,software.`softpath`,environmentals.`machine`,environmentals.`user`,environmentals.`passwd`,environmentals.`hostport`  FROM software INNER JOIN environmentals WHERE software.`domain` = environmentals.`domain`AND software.`Perject_Id` =? AND software.`Env_Id`=? AND software.`domain`=? AND environmentals.`machine`=?",PerjectId,EnvId,Domain,Machine).QueryRows(&list)
	//查询数据
	if  err == nil && qs>0{
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}




/*
 * 数据库初始化方法
 *@auth heyibo
 *@date 2018-5-30
 */
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Software))
	//创建表
	//orm.RunSyncdb("default", false, true)
}
