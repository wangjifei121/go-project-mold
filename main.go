package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"

	"projectDemo/config"
	"projectDemo/global"
	"projectDemo/initialize"
	"projectDemo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	//从命令行获取参数
	cfg = flag.String("config", "", "")
)

func main() {
	//解析参数
	flag.Parse()

	// init config 初始化Viper
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	initialize.InitZap() //初始化zap日志库
	initialize.InitDB()  //初始化数据库
	initialize.InitRedis() //初始化redis
	defer global.ProjectDB.Close()

	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	middlewares := []gin.HandlerFunc{}
	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := checkServer(); err != nil {
			global.ProjectLog.Error("自检程序发生错误...", err)
		}
		global.ProjectLog.Info(">>>当前服务启动成功，通过服务自检.")
	}()
	port := viper.GetString("addr")
	global.ProjectLog.Info(fmt.Sprintf("开始监听http地址%s", port))
	global.ProjectLog.Info(http.ListenAndServe(port, g).Error())

}

func checkServer() error {
	max := viper.GetInt("max_check_count")
	for i := 0; i < max; i++ {
		//发送一个GET请求给 `/check/health`，验证服务器是否成功.
		url := viper.GetString("url") + "/check/health"
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep 1 second 继续重试.
		global.ProjectLog.Info(">>>等待路由，1秒后重试。")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到路由. end")
}
