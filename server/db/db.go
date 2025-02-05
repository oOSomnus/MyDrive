package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	db   *sql.DB
	once sync.Once
}

var instance *DBManager
var once sync.Once

func GetDBManager(host, port, dbname, user, password string) *DBManager {
	once.Do(
		func() {
			fmt.Println("Initializing database connection...")
			instance = &DBManager{}
			instance.initialize(host, port, dbname, user, password)
		},
	)
	return instance
}

func (manager *DBManager) initialize(host, port, dbname, user, password string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	manager.db = db
	fmt.Println("Database connection initialized!")
}

func (manager *DBManager) GetDB() *sql.DB {
	return manager.db
}

func (manager *DBManager) Close() {
	if manager.db != nil {
		manager.db.Close()
	}
}
