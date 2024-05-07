package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ConnectionMySQL 连接 MySQL 数据库
func (a *App) ConnectionMySQL(dbtype string, username string, password string, dbname string) string {

	// 拼接连接信息
	var connectionInfo = username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname
	// 连接数据库
	db, err := sql.Open(dbtype, connectionInfo)

	if err != nil {
		panic(err)
	}

	// 验证连接是否正确
	err = db.Ping()
	if err != nil {
		return fmt.Sprintf("Errer connected to %s!", dbtype)
	}

	return fmt.Sprintf("Successfully connected to %s!", dbtype)
}
