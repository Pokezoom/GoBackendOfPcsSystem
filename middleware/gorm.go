package middleware

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var ErrNoInit = errors.New("gorm: please initialize with InitGorm() method")

var (
	isInit       = true                         // 是否初始化
	gormPool     = make(map[string]*gorm.DB, 5) // 连接单例池
	gormPoolLock sync.RWMutex                   // 连接单例池锁
)

// 需要自行配置
var myconfig = config{
	user:   "root",
	pass:   "990813",
	adrr:   "127.0.0.1",
	port:   "3306",
	dbname: "test",
}

type EGorm struct {
	MysqlConfName string
}

// MultipleConf 读写分离配置
type MultipleConf struct {
	Sources  []Conf // 写库
	Replicas []Conf // 读库
}

// Conf 单个数据库配置
type Conf struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	Params   string
}

// GDB 获取Gorm DB连接
// @return *gorm.DB
func (ctl *EGorm) GDB() *gorm.DB {
	return GetGorm(ctl.MysqlConfName)
}

// GetGorm 获取gorm连接
// @param name 配置名
// @return *gorm.DB
func GetGorm(name string) *gorm.DB {
	if !isInit {
		panic(ErrNoInit)
	}

	gormPoolLock.RLock()
	if db, ok := gormPool[name]; ok {
		gormPoolLock.RUnlock()
		return db
	}
	gormPoolLock.RUnlock()

	db, err := NewGormConnect(myconfig)
	if err != nil {
		panic("连接数据库失败")
		return &gorm.DB{}
	}

	gormPoolLock.Lock()
	gormPool[name] = db
	gormPoolLock.Unlock()

	return db
}

type config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbname string
}

// NewGormConnect 获取新客户端
// @param conf 配置信息
// @return *gorm.DB gorm连接
// @return error
func NewGormConnect(conf config) (*gorm.DB, error) {
	if !isInit {
		panic(ErrNoInit)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.user, conf.pass, conf.adrr, conf.port, conf.dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db, nil
}

//const dbDSNFormat = "%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local"
//
//func (c *Conf) DSN() string {
//	dsn := fmt.Sprintf(dbDSNFormat, c.User, c.Password, c.Host, c.Port, c.Database)
//	if c.Params != "" {
//		dsn = fmt.Sprintf("%s&%s", dsn, c.Params)
//	}
//	return dsn
//}
