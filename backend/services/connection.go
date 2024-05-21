package services

import (
	"ConchDBM/backend/utils/result"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

// SQLConnectObject 定义数据库连接对象
type SQLConnectObject struct {
	Driver   string `json:"driver"`   // 数据库驱动类型
	Host     string `json:"host"`     // 数据库地址
	Port     string `json:"port"`     // 数据库端口
	Username string `json:"username"` // 数据库用户名
	Password string `json:"password"` // 数据库密码
	Database string `json:"database"` // 数据库名
}

type Column struct {
	Key   string `json:"key"`
	Title string `json:"title"`
}

type Data struct {
	Data []map[string]interface{} `json:"data"`
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
		return result.Error(fmt.Sprintf("%s!", err))
	}

	// 拼接连接信息
	var connectionInfo = sqlObj.Username + ":" + sqlObj.Password + "@tcp(" + sqlObj.Host + ":" + sqlObj.Port + ")/" + sqlObj.Database

	// 连接数据库
	db, err := sql.Open(sqlObj.Driver, connectionInfo)

	if err != nil {
		return result.Error(fmt.Sprintf("%s!", err))
	}

	// 验证连接是否正确
	err = db.Ping()
	if err != nil {
		return result.Error(fmt.Sprintf("%s!", err))
	}

	dbConnection = db

	return result.Success("连接成功！", nil)
}

// Execute 执行SQL
func (c *Connection) Execute(sql string) string {

	rows, err := dbConnection.Query(sql)

	if err != nil {
		return result.Error(fmt.Sprintf("%s!", err))
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return result.Error(fmt.Sprintf("%s!", err))
	}

	columnObjs := make([]Column, len(columns))
	for i, column := range columns {
		columnObjs[i] = Column{Key: column, Title: column}
	}

	dataObjs := make([]map[string]interface{}, 0)

	for rows.Next() {
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		err := rows.Scan(columnPointers...)
		if err != nil {
			return result.Error(fmt.Sprintf("%s!", err))
		}

		dataObj := make(map[string]interface{})
		for i, columnValue := range columnValues {
			switch v := columnValue.(type) {
			case int64:
				dataObj[columns[i]] = strconv.Itoa(int(v))
			case []uint8:
				dataObj[columns[i]] = string(v)
			}
		}
		dataObjs = append(dataObjs, dataObj)
	}

	resultData := map[string]interface{}{
		"columns": columnObjs,
		"data":    Data{Data: dataObjs},
	}

	return result.Success("请求成功!", resultData)
}
