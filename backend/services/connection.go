package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义数据库连接对象
type SQLConnectObject struct {
	Driver   string `json:"driver"`   // 数据库驱动类型
	Host     string `json:"host"`     // 数据库地址
	Port     string `json:"port"`     // 数据库端口
	Username string `json:"username"` // 数据库用户名
	Password string `json:"password"` // 数据库密码
	Database string `json:"database"` // 数据库名
}

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
func (c *Connection) Connection(sqlObjStr string) string {
	var sqlObj SQLConnectObject

	err := json.Unmarshal([]byte(sqlObjStr), &sqlObj)

	if err != nil {
		return fmt.Sprintf("%s!", err)
	}

	// 拼接连接信息
	var connectionInfo = sqlObj.Username + ":" + sqlObj.Password + "@tcp(" + sqlObj.Host + ":" + sqlObj.Port + ")/" + sqlObj.Database

	// 连接数据库
	db, err := sql.Open(sqlObj.Driver, connectionInfo)

	if err != nil {
		return fmt.Sprintf("%s!", err)
	}

	// 验证连接是否正确
	err = db.Ping()
	if err != nil {
		return fmt.Sprintf("%s!", err)
	}
	dbConnection = db
	return fmt.Sprintf("Successfully connected to %s!", sqlObj.Driver)
}

// Execute 执行SQL
func (c *Connection) Execute(sql string) string {
	dbConnection.Query("123")
	return "项目重构ing..."
}
