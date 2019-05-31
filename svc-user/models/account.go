package models

import (
	log "github.com/sirupsen/logrus"
	db "leo/go-microservice/svc-user/databases"
)

type Account struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *Account) CreateAccount() (id int64, err error) {
	tx, _ := db.SqlDB.Begin()
	stm, err := tx.Prepare("insert into account(first_name, last_name) values (?, ?)")
	if err != nil {
		log.Error("can not create prepare statement")
		return
	}
	result, err := stm.Exec(a.FirstName, a.LastName)
	if err != nil {
		log.Error(err)
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		log.Error("commit error")
		rollErr := tx.Rollback()
		if rollErr != nil {
			log.Error("rollback error")
		}
	}
	id, err = result.LastInsertId()
	return
}
