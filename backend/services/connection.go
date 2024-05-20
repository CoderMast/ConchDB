package services

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connection struct
type Connection struct {
	ctx context.Context
}

// NewConnection creates a new Connection application struct
func NewConnection() *Connection {
	return &Connection{}
}

// Startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (c *Connection) Startup(ctx context.Context) {
	c.ctx = ctx
}

// 数据库连接指针
var dbConnection *sql.DB

// Connection 连接 MySQL 数据库
func (c *Connection) Connection(dbtype string, host string, port string, username string, password string, dbname string) string {

	// 拼接连接信息
	var connectionInfo = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	// 连接数据库
	db, err := sql.Open(dbtype, connectionInfo)

	if err != nil {
		return fmt.Sprintf("%s!", err)
	}

	// 验证连接是否正确
	err = db.Ping()
	if err != nil {
		return fmt.Sprintf("%s!", err)
	}
	dbConnection = db
	return fmt.Sprintf("Successfully connected to %s!", dbtype)
}

// Execute 执行SQL
func (c *Connection) Execute(sql string) string {
	dbConnection.Query("123")
	return "项目重构ing..."
}
