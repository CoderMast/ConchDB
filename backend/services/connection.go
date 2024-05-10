package services

import (
	"context"
	"database/sql"
	"encoding/json"
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
func (c *Connection) Connection(dbtype string, username string, password string, dbname string) string {

	// 拼接连接信息
	var connectionInfo = username + ":" + password + "@tcp(127.0.0.1:3306)/" + dbname
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
	if dbConnection == nil {
		return fmt.Sprintf("Error:请先连接数据库!")
	}

	rows, err := dbConnection.Query(sql)

	if err != nil {
		return fmt.Sprintf("Error:%s!", err)
	}
	// 确保关闭
	defer rows.Close()

	// 创建一个 Map 数组，用来存储返回数据
	results := make([]map[string]string, 0)

	// 获取列名
	columns, err := rows.Columns()

	if err != nil {
		return fmt.Sprintf("Error:%s!", err)
	}

	// 将列名加进返回切片对象
	columnMap := make(map[string]string)
	for _, column := range columns {
		columnMap[column] = column
	}
	results = append(results, columnMap)

	// 将内容添加进返回切片对象
	for rows.Next() {
		// 创建一个与列数量相同切片
		values := make([]string, len(columns))

		// 对应的指针切片
		ptrs := make([]interface{}, len(columns))

		for i := range values {
			ptrs[i] = &values[i]
		}
		// 这里实际上传入的事 value 中前 len(columns) 列的地址，故从 value 中能取到最后的实际数据
		err := rows.Scan(ptrs...)

		if err != nil {
			return fmt.Sprintf("Error:%s!", err)
		}

		// 创建一个值的 map 切片
		valuesMap := make(map[string]string)

		for i, value := range values {
			valuesMap[columns[i]] = value
		}

		results = append(results, valuesMap)
	}

	// 将返回对象转换为 JSON
	jsonResults, err := json.Marshal(results)

	if err != nil {
		return fmt.Sprintf("Error:%s!", err)
	}
	return fmt.Sprintf("%s", jsonResults)
}
