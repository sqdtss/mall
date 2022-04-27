package config

// 组合全部配置模型
type Config struct {
	Server  Server  `mapstructure:"server"`
	Mysql   Mysql   `mapstructure:"mysql"`
	Redis   Redis   `mapstructure:"redis"`
	Elastic Elastic `mapstructure:"elastic"`
	Upload  Upload  `mapstructure:"upload"`
	Jwt     Jwt     `mapstructure:"jwt"`
}

// 服务启动端口号配置
type Server struct {
	Port string `mapstructure:"port"`
}

// MySQL数据源配置
type Mysql struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Url             string `mapstructure:"url"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

// Redis配置
type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// Elastic配置
type Elastic struct {
	Url string `mapstructure:"url"`
}

// 文件上传相关路径配置
type Upload struct {
	SavePath  string `mapstructure:"savePath"`
	AccessUrl string `mapstructure:"accessUrl"`
}

// 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}
