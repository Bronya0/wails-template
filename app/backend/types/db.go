package types

// DBConfig 存放数据库所需的所有信息
type DBConfig struct {
	User       string `json:"user" yaml:"user"`
	Password   string `json:"password" yaml:"password"`
	SchemaName string `json:"schema_name" yaml:"schema_name"`
	DbAddress  string `json:"db_address" yaml:"db_address"` // 在SSH服务器看来，数据库的地址 "ip:port"
}
