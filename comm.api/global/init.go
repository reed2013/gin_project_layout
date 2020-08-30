package global

import (
	"comm.api/configs"
	"comm.pkgs/db"
	"comm.pkgs/redis"
	"comm.pkgs/zlog"
	"log"
)

//const BASE_DIR  = "D:/gocode/src/gin_project_layout/comm.api/"
const LogDir = "D:/gocode/src/gin_project_layout/comm.api/storage/logs/"

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