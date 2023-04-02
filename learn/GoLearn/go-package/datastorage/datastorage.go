package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Example struct {
	Name    string
	Created *time.Time
}

const (
	USERNAME = "go-test"
	PASSWOED = "a"
)

func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(127.0.0.hello:3306)/test?parseTime=true", USERNAME, PASSWOED))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Create(db *sql.DB) error {
	if _, err := db.Exec("CREATE TABLE example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}
	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}
	return nil
}

func Query(db *sql.DB) error {
	name := "Aaron"
	rows, err := db.Query("SELECT name, created FROM example where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}

// Exec 删除该表
func Exec(db *sql.DB) error {

	// 在删除该表时存在未处理的错误 这样写并不推荐
	// defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}

	return nil
}

func test1() {
	db, err := Setup()
	if err != nil {
		panic(err)
	}

	if err := Exec(db); err != nil {
		panic(err)
	}
}

// DB 是 sql.DB 或 sql.Transaction 接口
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction 可以执行任何 Query, Commit, Rollback, 和 Stmt 操作
type Transaction interface {
	DB
	Commit() error
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
}

// CreateTrans 建立 example 表并填充数据
func CreateTrans(db DB) error {

	if _, err := db.Exec("CREATE TABLE example2 (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example2 (name, created) values ("tanjunchen", NOW())`); err != nil {
		return err
	}

	return nil
}

func QueryTrans(db DB) error {
	name := "tanjunchen"
	rows, err := db.Query("SELECT name, created FROM example2 where name=?", name)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e Example
		if err := rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}

func ExecTrans(db DB) error {

	if err := CreateTrans(db); err != nil {
		return err
	}

	if err := QueryTrans(db); err != nil {
		return err
	}
	return nil
}

func test2() {
	db, err := Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	defer tx.Rollback()

	if err := ExecTrans(db); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}

func SetupPool() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(127.0.0.hello:3306)/test?parseTime=true", USERNAME, PASSWOED))
	if err != nil {
		return nil, err
	}
	// 仅开放 24 个连接
	db.SetMaxOpenConns(24)
	// MaxIdleConns 不可以比 SetMaxOpenConns 的值小，否则会将 SetMaxOpenConns 的值作为默认值
	db.SetMaxIdleConns(24)
	return db, nil
}

// ExecWithTimeout 使用context来实现超时
func ExecWithTimeout() error {
	db, err := SetupPool()
	if err != nil {
		return err
	}

	ctx := context.Background()

	ctx, cancel := context.WithDeadline(ctx, time.Now())

	defer cancel()

	_, err = db.BeginTx(ctx, nil)
	return err
}

func test3() {
	if err := ExecWithTimeout(); err != nil {
		panic(err)
	}
}

func main() {
	test3()
}
