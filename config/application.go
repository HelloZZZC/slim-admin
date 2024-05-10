package config

type Application struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
}
