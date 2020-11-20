package dbgo

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Config struct {
	User         string
	Password     string
	Host         string
	Port         string
	DB           string
	MaxIdleConns int `mapstructure:"max_idle_conns"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	Log          bool
}

func NewMysqlDB(dbConfig *Config) (*gorm.DB, error) {
	var err error
	Db, err := createConnection(dbConfig)
	if err != nil {
		return nil, err
	}
	if dbConfig.Log {
		Db.LogMode(true)
	}
	//Db.SetLogger(gormzap.New(logger.GetLogger()))
	return Db, nil
}

func createConnection(dbConfig *Config) (*gorm.DB, error) {
	var err error
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=UTF8&parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DB)
	db, err := gorm.Open("mysql", url)
	if err == nil {
		if dbConfig.MaxIdleConns != 0 && dbConfig.MaxOpenConns != 0 {
			db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
			db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
		}
	}
	return db, nil
}
