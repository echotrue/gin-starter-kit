package config

// Main config struct
type Config struct {
	Title    string `toml:"title"`
	Database *Database
	Redis    *Redis
	GinModel string `toml:"gin_model"`
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
