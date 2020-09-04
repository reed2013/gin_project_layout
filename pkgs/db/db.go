package db

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/reed2013/gin_project_layout/pkgs/config"
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

type Conf struct {
	UserName string
	Password string
	Host string
	DbName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

type DB struct {
	MasterDB *gorm.DB
	SlaveDB []*gorm.DB
}

var SingletonDB *DB
var once sync.Once
var masterConf Conf
var slaveConf []Conf
var masterConnErr error
var slaveConnErrs []error

func NewDB() error {
	once.Do(func(){
		var err error
		SingletonDB = &DB{}
		SingletonDB.MasterDB , err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			masterConf.UserName,
			masterConf.Password,
			masterConf.Host,
			masterConf.DbName,
			masterConf.Charset,
			masterConf.ParseTime,
		))
		if nil != err {
			masterConnErr = err
		}

		slaveSize := len(slaveConf)
		SingletonDB.SlaveDB = make([]*gorm.DB, slaveSize)
		slaveConnErrs = make([]error, slaveSize)
		if slaveSize > 0 {
			for i := 0; i < slaveSize; i++ {
				SingletonDB.SlaveDB[i], err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
					slaveConf[i].UserName,
					slaveConf[i].Password,
					slaveConf[i].Host,
					slaveConf[i].DbName,
					slaveConf[i].Charset,
					slaveConf[i].ParseTime,
				))
				if err != nil {
					slaveConnErrs[i] = err
				}
			}
		}
	})

	if masterConnErr != nil {
		return errors.New("master db conn error")
	}
	if len(slaveConnErrs) > 1 {
		return errors.New("slave db conn error")
	}

	return nil
}

func InitConf() error {
	var err error
	var ConfigManager *config.Config
	var configPath string
	if configPath, err = filepath.Abs("../../pkgs/db/"); err != nil {
		log.Fatal(err)
	}
	if ConfigManager, err = config.NewConfig(configPath, "config", "yaml"); err != nil {
		log.Fatal(err)
		return err
	}
	if err := ConfigManager.ReadSegment("Master", &masterConf); err != nil {
		return err
	}
	log.Println("db master conf: ", masterConf)
	if err := ConfigManager.ReadSegment("Slaves", &slaveConf); err != nil {
		return err
	}
	log.Println("db slave conf: ", slaveConf)

	return nil
}

func Master() *gorm.DB {
	if _, ok := interface{}(SingletonDB.MasterDB).(*gorm.DB); ok {
		return SingletonDB.MasterDB
	}

	return nil
}

func Slave() *gorm.DB {
	slaveSize := len(slaveConf)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(slaveSize)

	if _, ok := interface{}(SingletonDB.SlaveDB[index]).(*gorm.DB); ok {
		return SingletonDB.SlaveDB[index]
	}

	return nil
}

