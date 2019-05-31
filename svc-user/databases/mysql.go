package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

var log = logrus.New()
var SqlDB *sql.DB

func init() {
	var err error
	for i := 0; i < 3; i++ {
		SqlDB, err = connect()
		if err == nil {
			log.Info("connected to database")
			break
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func connect() (*sql.DB, error) {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/account?parseTime=true")
	if err != nil {
		log.WithFields(logrus.Fields{}).Error("could not connect to database")
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Error(err.Error())
	}
	return SqlDB, err
}

func DisConnect() {
	if SqlDB.Close() != nil {
		log.Error("database failed to disconnect")
	}
	log.Info("disconnected to database")
}
