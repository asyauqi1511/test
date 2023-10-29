package entity

type AppConfig struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	Hostname     string `yaml:"hostname"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
}
