package global

import (
	"github.com/reed2013/gin_project_layout/configs"
	"github.com/reed2013/gin_project_layout/pkgs/db"
	"github.com/reed2013/gin_project_layout/pkgs/redis"
	"github.com/reed2013/gin_project_layout/pkgs/zlog"
	"log"
)

//const BASE_DIR  = "D:/gocode/src/gin_project_layout/github.com/reed2013/gin_project_layout/"
const LogDir = "D:/gocode/src/gin_project_layout/github.com/reed2013/gin_project_layout/storage/logs/"

var (
	ServerConf *configs.ServerConf
	AppConf *configs.AppConf
)

func Init() {
	zlog.NewLogger(LogDir + "api.log")
	if err := InitConfig(); err != nil {
		zlog.SugarLogger().Error(err)
	}
	if err := db.NewDB(); err != nil {
		zlog.SugarLogger().Error(err)
	}
	if err := redis.NewRedis(); err != nil {
		zlog.SugarLogger().Error(err)
	}

}

func InitConfig() error {
	config, err := configs.NewConfig()
	if  err != nil {
		return err
	}
	if err := config.ReadSegment("Server", &ServerConf); err != nil {
		return err
	}
	log.Println("server conf: ", ServerConf)
	if err := config.ReadSegment("App", &AppConf); err != nil {
		return err
	}
	log.Println("server conf: ", AppConf)

	return nil
}