package global

import (
	"comm.api/configs"
	"comm.pkgs/db"
	"comm.pkgs/redis"
	"log"
)

var (
	ServerConf *configs.ServerConf
	AppConf *configs.AppConf
)

func Init() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}
	if err := InitDb(); err != nil {
		log.Fatal(err)
	}
	/*
	if err := InitRedis(); err != nil {
		log.Fatal(err)
	}
	*/
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

func InitDb() error {
	if err := db.NewDB(); err != nil {
		return err
	}

	return nil
}

func InitRedis() error {
	if err := redis.NewRedis(); err != nil {
		return err
	}

	return nil
}