package redis

import (
	"github.com/reed2013/gin_project_layout/pkgs/config"
	"context"
	"errors"
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
var masterConnErr error
var slaveConnErrs []error

type Conf struct {
	Addr string
	DB int
	Password string
}

type Redis struct {
	MasterRedis *redis.Client
	SlaveRedis []*redis.Client
}

func NewRedis() error {
	once.Do(func(){
		SingletonRedis = &Redis{}
		SingletonRedis.MasterRedis = redis.NewClient(&redis.Options{
			Addr:     masterConf.Addr,
			Password: masterConf.Password,
			DB:       masterConf.DB,
		})
		_, err := SingletonRedis.MasterRedis.Ping(ctx).Result()
		if err != nil {
			masterConnErr = err
		}

		slaveSize := len(slaveConf)
		SingletonRedis.SlaveRedis = make([]*redis.Client, slaveSize)
		slaveConnErrs = make([]error, slaveSize)
		if slaveSize > 0 {
			for i := 0; i < slaveSize; i++ {
				SingletonRedis.SlaveRedis[i] = redis.NewClient(&redis.Options{
					Addr:     slaveConf[i].Addr,
					Password: slaveConf[i].Password,
					DB:       slaveConf[i].DB,
				})
				_, err := SingletonRedis.SlaveRedis[i].Ping(ctx).Result()
				if err != nil {
					slaveConnErrs[i] = err
				}
			}
		}
	})

	if masterConnErr != nil {
		return errors.New("master redis conn error")
	}
	if len(slaveConnErrs) > 1 {
		return errors.New("slave redis conn error")
	}

	return nil
}

func InitConf() error {
	var err error
	var ConfigManager *config.Config
	var configPath string
	if configPath, err = filepath.Abs("../../pkgs/redis/"); err != nil {
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

	return SingletonRedis.SlaveRedis[index]
}