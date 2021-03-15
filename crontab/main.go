package main

import (
    "fmt"
    "time"
    "io/ioutil"
    "encoding/json"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

const(
    ConfFile = "/home/yaoxiaofeng/test/server.json"
)

// 配置
// --------------------------------------------------------
func GetConfig(file string) *Config {
    conf := &Config{}
    content, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    json.Unmarshal(content, conf)
    return conf
}

type Config struct {
    Server struct {
        Host string
        Port int
    }

    Db struct {
        Host   string
        Port   int
        User   string
        Passwd string
        DbName string
    }

    Redis struct {
        Host     string
        Port     int
        Database int
    }
}

func main() {
	var ch chan int
    ticker := time.NewTicker(time.Second * 5)
    go func() {
        for range ticker.C {
            handle()
        }
    }()
	<-ch
}

func handle() {
    // 加载配置
    conf := GetConfig(ConfFile)

    // 链接数据库
    dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        conf.Db.User,
        conf.Db.Passwd,
        conf.Db.Host,
        conf.Db.DbName,
    )
    db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

    // 执行SQL
    db.Exec("UPDATE t_user SET balance=? WHERE id=46399747 limit 1", time.Now().Format("150405"))
}

