package config

// Main config struct
type Config struct {
	Title    string `toml:"title"`
	Database *Database
	Redis    *Redis
	GinModel string `toml:"gin_model"`
	Log      *Log
}

// Database configuration
type Database struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Db       string `toml:"db"`
}

// Redis configuration
type Redis struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
	Auth string `toml:"auth"`
}

type Log struct {
	Target  string `toml:"target"`
	Level   string `toml:"level"`
	File    *LogFile
	Network *LogNetwork
}


type LogFile struct {
	Filepath string `toml:"filepath"`
}
type LogNetwork struct {
	Url string `toml:"url"`
}
