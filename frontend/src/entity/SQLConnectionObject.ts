// 数据库连接对象
export interface SQLConnectObject {
  driver: string   // 数据库驱动类型
  host: string     // 数据库地址
  port: string     // 数据库端口
  username: string // 数据库用户名
  password: string // 数据库密码
  database: string // 数据库名
}