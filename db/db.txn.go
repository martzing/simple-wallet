package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/martzing/simple-wallet/configs"
	"gorm.io/gorm"
)

type DatabaseTransaction struct {
	db            *gorm.DB
	transactionId string
}

var REPEATABLE_READ = "REPEATABLE READ"
var READ_COMMITTED = "READ COMMITTED"
var READ_UNCOMMITTED = "READ UNCOMMITTED"
var SERIALIZABLE = "SERIALIZABLE"

func NewTransaction() DatabaseTransaction {
	dbTxn := DatabaseTransaction{}
	return dbTxn
}

func executeTransaction(db *gorm.DB, transactionId string) {
	statement := db.Statement
	sql := db.Explain(statement.SQL.String(), statement.Vars...)

	fmt.Printf("Transaction (%s): %s\n", transactionId, sql)
}

func registerCallback(db *gorm.DB, transactionId string, isolationLevel string) {
	isolationLevelSql := fmt.Sprintf("SET TRANSACTION ISOLATION LEVEL %s", isolationLevel)
	db = db.Exec(isolationLevelSql)
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

func (dbTxn *DatabaseTransaction) Begin(isolationLevel string) {
	transactionId := uuid.New().String()

	Connect(*configs.DbConfig)

	registerCallback(DB, transactionId, isolationLevel)

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
