package redis

import (
	"comm.pkgs/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"math/rand"
	"path/filepath"
	"sync"
	"time"
)

func init() {
	if err := InitConf(); nil != err {
		log.Fatal(err)
	}
}

var once sync.Once
var ctx = context.Background()
var masterConf Conf
var slaveConf []Conf
var SingletonRedis *Redis

type Conf struct {
	Addr string
	DB int
	Password string
}

type Redis struct {
	MasterRedis *redis.Client
	SlaveRedis *redis.Client
}

func NewRedis() error {
	once.Do(func(){
		SingletonRedis = &Redis{}
	})

	SingletonRedis.MasterRedis = redis.NewClient(&redis.Options{
		Addr:     masterConf.Addr,
		Password: masterConf.Password,
		DB:       masterConf.DB,
	})
	pong, err := SingletonRedis.MasterRedis.Ping(ctx).Result()
	fmt.Println(pong, err)

	return nil
}

func InitConf() error {
	var err error
	var ConfigManager *config.Config
	var configPath string
	if configPath, err = filepath.Abs("../../comm.pkgs/redis/"); err != nil {
		log.Fatal(err)
	}
	if ConfigManager, err = config.NewConfig(configPath, "config", "yaml"); err != nil {
		log.Fatal(err)
		return err
	}
	if err := ConfigManager.ReadSegment("Master", &masterConf); err != nil {
		return err
	}
	log.Println("redis master conf: ", masterConf)
	if err := ConfigManager.ReadSegment("Slaves", &slaveConf); err != nil {
		return err
	}
	log.Println("redis slave conf: ", slaveConf)

	return nil
}

func Master() *redis.Client {
	if _, ok := interface{}(SingletonRedis.MasterRedis).(*redis.Client); ok {
		return SingletonRedis.MasterRedis
	}

	return nil
}

func Slave() *redis.Client {
	slaveSize := len(slaveConf)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(slaveSize)

	SingletonRedis.SlaveRedis = redis.NewClient(&redis.Options{
		Addr:     slaveConf[index].Addr,
		Password: slaveConf[index].Password,
		DB:       slaveConf[index].DB,
	})

	return SingletonRedis.SlaveRedis
}