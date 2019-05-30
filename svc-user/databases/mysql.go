package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/account?parseTime=true")
	if err != nil {
		log.WithFields(logrus.Fields{}).Error("could not connect to database")
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Error(err.Error())
	}
}
