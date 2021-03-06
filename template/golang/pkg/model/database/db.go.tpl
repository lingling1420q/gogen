// Code generated by gogen.
// DO NOT EDIT!

package {{underline .Pkg}}

import (
	"fmt"

	"github.com/silenceper/gogen/pkg/orm"
	"github.com/jinzhu/gorm"
)

var (
	dbName = "{{.Database}}"

	dbConfig *orm.Options
	dbInfo   *orm.DBInfo
)

//InitDB 设置db链接
func InitDB(opts *orm.Options){
    dbConfig=opts
}

//GetDB 获取db连接信息
func GetDB() *orm.DBInfo {
	if dbInfo != nil {
		return dbInfo
	}
	if dbConfig == nil {
		panic("must init db")
	}

	if dbConfig.DataSource == "" {
		dbConfig.DataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.DBName)
	}

	database, err := gorm.Open("mysql", dbConfig.DataSource)
	if err != nil {
		panic(err)
	}
	database.LogMode(true)
	if dbConfig.MaxIdleConns != 0 {
		database.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
	}
	if dbConfig.MaxOpenConns != 0 {
		database.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
	}

	//try
	err = database.DB().Ping()
	if err != nil {
		panic(err)
	}

	//custom logger
	if dbConfig.EnableLog {
		database.LogMode(true)
	} else {
		database.LogMode(false)
	}
	database.SetLogger(&orm.GormLogger{})

	dbInfo = &orm.DBInfo{DB: database}
	return dbInfo
}

