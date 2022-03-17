package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"path"
	"projectDemo/global"
	"projectDemo/models"
)

func InitDB() *gorm.DB {
	var err error

	// 数据库类型：mysql/sqlite3
	dbType := viper.GetString("database.type")

	// mysql配置信息
	dbName := viper.GetString("database.name")         // 数据库名称
	dbHost := viper.GetString("database.host")         // 数据库ip地址
	dbPort := viper.GetString("database.port")         // 数据库端口
	dbUsername := viper.GetString("database.username") // 用户名
	dbPwd := viper.GetString("database.password")      // 密码
	dbCharset := viper.GetString("database.charset")   // 指定字符集

	var dataSource string
	switch dbType {
	case "mysql":
		// parseTime=true 表示自动解析为时间
		// db, _ = gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local")
		dataSource = dbUsername + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=true"
		global.ProjectDB, err = gorm.Open(dbType, dataSource)
	case "sqlite3":
		dataSource = "database" + string(os.PathSeparator) + dbName
		if !gfile.Exists(dataSource) {
			os.MkdirAll(path.Dir(dataSource), os.ModePerm)
			os.Create(dataSource)
		}
		global.ProjectDB, err = gorm.Open(dbType, dataSource)
	}

	if err != nil {
		global.ProjectLog.Error(">>>初始化数据库失败", err)
	}

	// 设置数据库操作显示原生SQL 语句
	global.ProjectDB.LogMode(true)

	// 表明禁用后缀加s
	global.ProjectDB.SingularTable(true)

	global.ProjectLog.Info(fmt.Sprintf(">>>>初始化%v数据库成功", dbType))

	Migrate(dbType)

	return global.ProjectDB
}

//数据库表迁移
func Migrate(dbType string) {
	is_migrate := viper.GetBool("migrate")
	if is_migrate {
		global.ProjectLog.Info(fmt.Sprintf(">>>正在进行%v数据库表迁移", dbType))
		// 数据库表动态迁移
		global.ProjectDB.AutoMigrate(&models.User{})
		global.ProjectLog.Info(fmt.Sprintf(">>>%v数据库表迁移完成", dbType))

	}
}
