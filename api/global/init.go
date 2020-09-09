package global

import (
	"github.com/opentracing/opentracing-go"
	"github.com/reed2013/gin_project_layout/api/configs"
	"github.com/reed2013/gin_project_layout/pkgs/db"
	"github.com/reed2013/gin_project_layout/pkgs/jtracer"
	"github.com/reed2013/gin_project_layout/pkgs/redis"
	"github.com/reed2013/gin_project_layout/pkgs/zlog"
	"go.opentelemetry.io/otel/api/global"
	"log"
)

//const BASE_DIR  = "D:/gocode/src/gin_project_layout/github.com/reed2013/gin_project_layout/"
const LogDir = "D:/gocode/src/gin_project_layout/api/storage/logs/"

var (
	ServerConf *configs.ServerConf
	AppConf *configs.AppConf
	Tracer opentracing.Tracer
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
	if err := InitTracer(); err != nil {
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

func InitTracer() error {
	var err error
	Tracer, _, err = jtracer.NewTracer("gin_project_layout_service", "127.0.0.1:6381")
	if err != nil {
		return err
	}

	return nil
}

