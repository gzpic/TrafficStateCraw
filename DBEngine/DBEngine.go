package DBEngine

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type databaseSetting struct {
	DBType    string
	UserName  string
	Password  string
	Host      string
	DBName    string
	Charset   string
	ParseTime string
}

func InitDBEngine() *gorm.DB {
	DBsetting := &databaseSetting{
		"mysql",
		"root",
		"1996",
		"localhost:3306",
		"testgo",
		"utf8",
		"True",
	}
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local"
	db, err := gorm.Open(DBsetting.DBType, fmt.Sprintf(s,
		DBsetting.UserName,
		DBsetting.Password,
		DBsetting.Host,
		DBsetting.DBName,
		DBsetting.Charset,
		DBsetting.ParseTime,
	))

	if err != nil {
		fmt.Println(err)
	}
	return db
}
