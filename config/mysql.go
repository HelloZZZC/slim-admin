package config

type Mysql struct {
	Url            string         `yaml:"url"`
	Username       string         `yaml:"username"`
	Password       string         `yaml:"password"`
	ConnectionPool ConnectionPool `yaml:"connection-pool"`
}

type ConnectionPool struct {
	MaxIdleConns           int `yaml:"max-idle-conns"`
	MaxOpenConns           int `yaml:"max-open-conns"`
	ConnMaxLifeTimeSeconds int `yaml:"conn-max-life-time-seconds"`
}
