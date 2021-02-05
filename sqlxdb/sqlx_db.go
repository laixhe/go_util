package sqlxdb

import (
	"database/sql"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/laixhe/goutil/zaplog"
)

var (
	once sync.Once

	db *sqlx.DB
)

// Init 初始化数据库
// dbType  数据库类型     [mysql]
// dsn     数据库连接信息  [root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local]
// openMax 设置最大打开的连接数
// idleMax 设置闲置的连接数
func Init(dbType, dsn string, openMax, idleMax int) {

	once.Do(func() {
		var err error
		db, err = sqlx.Connect(dbType, dsn)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		// 设置闲置的连接数
		db.SetMaxIdleConns(idleMax)
		// 设置最大打开的连接数
		db.SetMaxOpenConns(openMax)
	})

}

// DB get sqlx
func DB() *sqlx.DB {
	return db
}

// Get 获取一条 sql 数据
func Get(dest interface{}, query string, args ...interface{}) error {
	zaplog.Debugf("get query[%s] args[%v]", query, args)

	return db.Get(dest, query, args...)
}

// Select 获取多条 sql 数据
func Select(dest interface{}, query string, args ...interface{}) error {
	zaplog.Debugf("select query[%s] args[%v]", query, args)

	return db.Select(dest, query, args...)
}

// SelectIn 获取多条 sql 数据
func SelectIn(dest interface{}, query string, args ...interface{}) error {
	zaplog.Debugf("select in query[%s] args[%v]", query, args)

	queryIN, args, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	return db.Select(dest, queryIN, args...)
}

// Exec 执行 sql 语句 insert delete update
func Exec(query string, args ...interface{}) (sql.Result, error) {
	zaplog.Debugf("exec query[%s] args[%v]", query, args)

	return db.Exec(query, args...)
}

// ExecIn 执行 sql 语句 insert delete update
func ExecIn(query string, args ...interface{}) (sql.Result, error) {
	zaplog.Debugf("exec in query[%s] args[%v]", query, args)

	queryIN, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return db.Exec(queryIN, args...)
}

// DBTx 事务相关
type DBTx struct {
	db *sqlx.Tx
}

// Begin 开始事务
func Begin() (*DBTx, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	return &DBTx{db: tx}, nil
}

// Commit 提交事务
func (tx *DBTx) Commit() error {
	return tx.db.Commit()
}

// Rollback 回滚事务
func (tx *DBTx) Rollback() error {
	return tx.db.Rollback()
}

// Get 获取一条 sql 数据
func (tx *DBTx) Get(dest interface{}, query string, args ...interface{}) error {
	zaplog.Debugf("get tx query[%s] args[%v]", query, args)

	return tx.db.Get(dest, query, args...)
}

// Select 获取多条 sql 数据
func (tx *DBTx) Select(dest interface{}, query string, args ...interface{}) error {
	zaplog.Debugf("select tx query[%s] args[%v]", query, args)

	return tx.db.Select(dest, query, args...)
}

// SelectIn 获取多条 sql 数据
func (tx *DBTx) SelectIn(dest interface{}, query string, args ...interface{}) error {
	zaplog.Debugf("select in tx query[%s] args[%v]", query, args)

	queryIN, args, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}

	return tx.db.Select(dest, queryIN, args...)
}

// Exec 执行 sql 语句 insert delete update
func (tx *DBTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	zaplog.Debugf("exec tx query[%s] args[%v]", query, args)

	return tx.db.Exec(query, args...)
}

// ExecIn 执行 sql 语句 insert delete update
func (tx *DBTx) ExecIn(query string, args ...interface{}) (sql.Result, error) {
	zaplog.Debugf("exec in tx query[%s] args[%v]", query, args)

	queryIN, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	return tx.db.Exec(queryIN, args...)
}
