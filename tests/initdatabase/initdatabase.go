package initdatabase

import (
	"fmt"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMyDB(t *testing.T) {
	//注册数据库服务
	beego.Debug("正在初始化数据库连接信息...")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		t.Fatal("数据库链接失败，失败信息为：", err)
		os.Exit(1)
	}
	dburl := "%s:%s@tcp(%s:%d)/%s?charset=%s"
	dburl = fmt.Sprintf(dburl,
		"root",
		"123456",
		"192.168.10.134",
		3306,
		"logistics_test",
		"utf8")
	beego.Informational(fmt.Sprintf("数据库连接信息为： %s", dburl))
	err = orm.RegisterDataBase("default", "mysql", dburl, 5, 20)
	if err != nil {
		t.Fatal("数据库链接失败，失败信息为：", err)
		os.Exit(1)
	}
	orm.Debug = true
}
