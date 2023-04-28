package db

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DatabaseTransaction struct {
	db            *gorm.DB
	transactionId string
}

func NewTransaction() DatabaseTransaction {
	dbTxn := DatabaseTransaction{}
	return dbTxn
}

func executeTransaction(db *gorm.DB, transactionId string) {
	statement := db.Statement
	sql := db.Explain(statement.SQL.String(), statement.Vars...)

	fmt.Printf("Transaction (%s): %s\n", transactionId, sql)
}

func registerCallback(db *gorm.DB, transactionId string) {
	callback := db.Callback()

	callback.Create().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Query().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Update().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Delete().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Raw().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Row().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
}

func (dbTxn *DatabaseTransaction) Begin() {
	transactionId := uuid.New().String()

	config := DBConfig{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "root",
		DBName:   "db",
	}

	Connect(config)

	registerCallback(DB, transactionId)

	dbTxn.db = DB.Begin()
	dbTxn.transactionId = transactionId

	fmt.Printf("Transaction (%s): Start Transaction\n", dbTxn.transactionId)
}

func (dbTxn *DatabaseTransaction) Get() *gorm.DB {
	return dbTxn.db
}

func (dbTxn *DatabaseTransaction) Commit() {
	if dbTxn.db == nil {
		panic("Database transaction already commit or rollback")
	}
	dbTxn.db.Commit()
	dbTxn.db = nil
	fmt.Printf("Transaction (%s): Commit\n", dbTxn.transactionId)
}

func (dbTxn *DatabaseTransaction) Rollback() {
	if dbTxn.db == nil {
		panic("Database transaction already commit or rollback")
	}
	dbTxn.db.Rollback()
	dbTxn.db = nil
	fmt.Printf("Transaction (%s): Rollback\n", dbTxn.transactionId)
}
