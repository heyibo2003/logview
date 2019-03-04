package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type Environmentals struct {
	Id                   int      `orm:"pk:auto"`
	Name                 string   `orm:"size(255)"`
	Path                 string   `orm:"size(255)"`
	Machine              string   `orm:"size(255)"`
	User                 string   `orm:"size(255)"`
	Passwd               string   `orm:"size(255)"`
	Hostport             int      `orm:"size(11)"`
	EnvId               int      `orm:"size(10)"`
	Domain               string   `orm:"size(255)"`
	PerjectId           int      `orm:"size(10)"`
	CreateTime           time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdateTime           time.Time   `orm:"auto_now;type(datetime)"`
}

type options struct {
	Name    string
	Path    string
	Machine string
	User    string
	Passwd  string
	Hostport    int
}

type applist struct {
	Id                   int
	Name                 string
	Path                 string
	Machine              string
	User                 string
	Passwd               string
	Hostport             int
	Domain               string
	Envname              string
	Perject              string
	CreateTime           time.Time
	UpdateTime           time.Time
}



func (this *Environmentals) TableName() string {
	return "environmentals"
}
//根据应用名称模糊查询
func (c *Environmentals) GetEnvironmentalsByDomain(Domain string) (dataList []*applist,err error ) {
	o := orm.NewOrm()
	o.Using("default")
	var list []*applist
	domain :=Domain
	result,err:= o.Raw("SELECT  environmentals.id,environmentals.name,environmentals.path,environmentals.machine,environmentals.user,environmentals.passwd,environmentals.hostport,environmentals.domain,environmentals.create_time,environmentals.update_time,env.envname,perject.perject FROM `environmentals` INNER JOIN env,perject WHERE environmentals.`Env_Id`=env.`id` AND environmentals.`Perject_Id`=perject.`id`  AND environmentals.`domain` LIKE '%"+domain+"%'").QueryRows(&list)
	//查询数据
	if  err == nil && result>0{
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}

//表查询整个表
func (c *Environmentals) GetAll () (dataList []*applist,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*applist
	result,err:= o.Raw("SELECT  environmentals.id,environmentals.name,environmentals.path,environmentals.machine,environmentals.user,environmentals.passwd,environmentals.hostport,environmentals.domain,environmentals.create_time,environmentals.update_time,env.envname,perject.perject FROM `environmentals` INNER JOIN env,perject WHERE environmentals.`Env_Id`=env.`id` AND environmentals.`Perject_Id`=perject.`id`").QueryRows(&list)
	//查询数据
	if  err == nil && result>0{
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}
//根据envID参数查询表environmentals的domain
func (c *Environmentals) GetDomainsByEnv(EnvId int)   (dataList []*Environmentals,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Environmentals
	qs := o.QueryTable("environmentals").Filter("EnvId",EnvId).GroupBy("domain")
	//查询数据
	if _, err = qs.All(&list,"domain"); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-domain :", dataList)
		return dataList, nil

	}
	logs.Info("dataList-daomain-2 :",dataList)
	return dataList, err
}
//根据envID和domain参数查询表environmentals的Machine字段
func (c * Environmentals) GetMachineBydDomain(EnvId int, domain string) (dataList []*Environmentals,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Environmentals
	qs := o.QueryTable("environmentals").Filter("EnvId",EnvId).Filter("domain",domain).GroupBy("machine")
	//查询数据
	if _, err = qs.All(&list,"machine"); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-machine :", dataList)
		return dataList, nil

	}
	logs.Info("dataList-machine-2 :",dataList)
	return dataList, err
}
//根据PerjectId,envID和domain参数查询表environmentals的Machine字段
func (c * Environmentals) GetMachineBydDomainAndEnvIdAndPerjectId(PerjectId int,EnvId int, domain string) (dataList []*Environmentals,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Environmentals
	qs := o.QueryTable("environmentals").Filter("PerjectId",PerjectId).Filter("EnvId",EnvId).Filter("domain",domain).GroupBy("machine")
	//查询数据
	if _, err = qs.All(&list,"machine"); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-machine :", dataList)
		return dataList, nil

	}
	logs.Info("dataList-machine-2 :",dataList)
	return dataList, err
}

//根据envID和domain、Machine参数查询表environmentals的name字段
func (c *Environmentals) GetNameByDomainAndMachine(EnvId int, domain string,machine string)  (dataList []*Environmentals,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Environmentals
	qs := o.QueryTable("environmentals").Filter("EnvId",EnvId).Filter("domain",domain).Filter("machine",machine).GroupBy("name")
	//查询数据
	if _, err = qs.All(&list,"name"); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-name :", dataList)
		return dataList, nil

	}
	logs.Info("dataList-name-2 :",dataList)
	return dataList, err
}

//根据envID和domain、Machine、name参数查询表environmentals

func (c *Environmentals) GetEnvironmentalsByFourParameter(EnvId int, domain string,machine string,name string) (dataList []*Environmentals,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Environmentals
	qs := o.QueryTable("environmentals").Filter("EnvId",EnvId).Filter("domain",domain).Filter("machine",machine).Filter("name",name)
	//查询数据
	if _, err = qs.All(&list,"path","user","passwd","hostport"); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-all :", dataList)
		return dataList, nil

	}
	logs.Info("dataList-all-2 :",dataList)
	return dataList, err
}

func (c *Environmentals) GetEnvironmentalsById(id int) (v *Environmentals, err error){
	o := orm.NewOrm()
	o.Using("default")
	u :=&Environmentals{Id: id}
	if err = o.Read(u, "Id"); err == nil {
		return u, nil
	}
	return nil, err
}

func (c *Environmentals) Add(Environmental *Environmentals) (id int,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Insert(Environmental); err == nil {
		logs.Info(" Insert in database:", result)
		return Environmental.Id, nil
	}
	return 0, err
}
func (c *Environmentals) Update(Environmental *Environmentals) (err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Update(Environmental); err == nil {
		logs.Info(" updated in database:", result)
	}
	return
}

func (c *Environmentals) Delete (id int) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	u:=Environmentals{Id:id}
	if err = o.Read(&u, "Id"); err == nil {
		if num, err := o.Delete(&Environmentals{Id: id}); err == nil {
			logs.Info(" deleted in usertable:", num)
		}
	}
	return
}





func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Environmentals))
	//创建表
	//orm.RunSyncdb("default", false, true)
}