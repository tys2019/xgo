package config

type Cfg struct {
	LogConfig  *LogConfig `json:"log_config"`
	
	DB DBConfig`json:"db"`
}



// 日志配置信息
type LogConfig struct {
	Count string `json:"count"`
	Level string `json:"level"`
	Uint  string `json:"uint"`
	RoleType string `json:"role_type"`
	File string `json:"file"`
	Dir string `json:"dir"`
	Size int `json:"size"`
	Compress int `json:"compress"`
}

// mysql 数据库相关
type DBConfig struct {
	Username string `json:"username"`
	Host string `json:"host"`
	Password  string `json:"password"`
	Port int `json:"port"`
	Databases string `json:"databases"`
}