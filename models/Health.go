package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"time"
	"fmt"
)
type Health struct {
	Id                   int      `orm:"pk:auto"`
	Name                 string   `orm:"size(100)"`
	Processname          string   `orm:"size(100)"`
	Inspection           string   `orm:"size(50)"`
	Url                  string   `orm:"size(255)"`
	Port                 int      `orm:"size(11)"`
	User                 string   `orm:"size(100)"`
	Passwd               string   `orm:"size(100)"`
	Machine              string   `orm:"size(100)"`
	Hostport             int      `orm:"size(11)"`
	EnvId               int      `orm:"size(10)"`
	PerjectId           int      `orm:"size(10)"`
	CreateTime           time.Time   `orm:"auto_now_add;type(datetime)"`
	UpdateTime           time.Time   `orm:"auto_now;type(datetime)"`
}
type PerjectHealth struct {
	Id                   int
	Name                 string
	Processname          string
	Inspection           string
	Url                  string
	Port                 int
	User                 string
	Passwd               string
	Machine              string
	Hostport             int
	Envname              string
	Perject              string
	CreateTime           time.Time
	UpdateTime           time.Time
}
func (this *Health) TableName() string {
	return "health"
}

//根据项目Id和环境Id查询
func (c *Health) GetHealthByPerjectIdAndEnvId(PerjectId int,EnvId int) (dataList []*Health,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*Health
	qs := o.QueryTable("health").Filter("PerjectId",PerjectId).Filter("EnvId",EnvId)
	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		logs.Info("dataList-name :", dataList)
		return dataList, nil

	}
	logs.Info("dataList-name-2 :",dataList)
	return dataList, err

}

/*
 * 根据服务名称模糊查询
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Health) GetPerjectHealthListByName (Name string) (dataList []*PerjectHealth,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*PerjectHealth
	qs,err:= o.Raw("SELECT health.`id`,health.`name`,health.`processname`,health.`inspection`,health.`url`,health.`port`,health.`user`,health.`passwd`,health.`machine`,health.`hostport`,perject.`perject`,env.`envname`,health.`Create_Time`,health.`Update_Time`FROM health INNER JOIN perject,env WHERE health.`Perject_Id`=perject.`id` AND health.`Env_Id`=env.`id` AND health.`name`LIKE '%"+Name+"%'").QueryRows(&list)
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
 * 查询所有用户方法
 *@auth heyibo
 *@date 2018-5-30
 */

func (c *Health) GetAll() (dataList []*PerjectHealth,err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	var list []*PerjectHealth
	result,err:= o.Raw("SELECT health.`id`,health.`name`,health.`processname`,health.`inspection`,health.`url`,health.`port`,health.`user`,health.`passwd`,health.`machine`,health.`hostport`,perject.`perject`,env.`envname`,health.`Create_Time`,health.`Update_Time`FROM health INNER JOIN perject,env WHERE health.`Perject_Id`=perject.`id` AND health.`Env_Id`=env.`id`").QueryRows(&list)
	//查询数据
	if  err == nil && result>0{
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return dataList, err
}

/*
 * 添加环境名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c *Health) Add(Health *Health) (id int,err error ){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if result, err := o.Insert(Health); err == nil {
		logs.Info(" Insert in database:", result)
		return Health.Id, nil
	}
	return Health.Id, err
}


/*
 * 修改环境名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c * Health) Update(Health *Health) (err error ) {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("ss:",Health)
	o.Update(Health)
	return
}


/*
 * 删除环境名称方法
 *@auth heyibo
 *@date 2018-5-30
 */
func (c * Health)  Delete(id int) (err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=Health{Id:id}
	if err = o.Read(&u, "Id"); err == nil {
		if num, err := o.Delete(&Health{Id: id}); err == nil {
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
func (c *Health) GetHealthById(id int) (v *Health, err error) {
	o := orm.NewOrm()
	o.Using("default")
	u :=&Health{Id: id}
	if err = o.Read(u, "Id"); err == nil {
		return u, nil
	}
	return nil, err
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Health))
	//创建表
	//orm.RunSyncdb("default", false, true)
}