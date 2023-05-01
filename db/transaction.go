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

var DB *gorm.DB

func ConnectDB() error {
	db, err := Connect(*configs.DbConfig)

	DB = db

	registerCallback(db, nil, nil)

	return err
}

func NewTransaction() DatabaseTransaction {
	dbTxn := DatabaseTransaction{}
	return dbTxn
}

func executeTransaction(db *gorm.DB, transactionId *string) {
	statement := db.Statement
	sql := db.Explain(statement.SQL.String(), statement.Vars...)

	tid := "default"
	if transactionId != nil {
		tid = *transactionId
	}
	fmt.Printf("Transaction (%s): %s\n", tid, sql)
}

func registerCallback(db *gorm.DB, transactionId *string, isolationLevel *string) {
	if isolationLevel != nil {
		isolationLevelSql := fmt.Sprintf("SET TRANSACTION ISOLATION LEVEL %s", *isolationLevel)
		db = db.Exec(isolationLevelSql)
	}
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

func (dbTxn *DatabaseTransaction) Begin(isolationLevel string) (*gorm.DB, error) {
	transactionId := uuid.New().String()

	db, err := Connect(*configs.DbConfig)

	if err != nil {
		return nil, err
	}

	registerCallback(db, &transactionId, &isolationLevel)

	dbTxn.db = db.Begin()
	dbTxn.transactionId = transactionId

	fmt.Printf("Transaction (%s): Start Transaction\n", dbTxn.transactionId)

	return db, nil
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
