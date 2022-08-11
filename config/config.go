package config

type Config struct {
	AppPort    string `json:"appport"`
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	DBHost     string `json:"dbhost"`
	DBPort     int    `json:"dbport"`
	DBName     string `json:"dbname"`
}

func (c *Config) LoadConfig() error {
	// Implement viper load config
	c.AppPort = ":8082"
	c.DBHost = "localhost"
	c.DBPort = 5432
	c.DBPassword = "password"
	c.DBUser = "postgres"
	c.DBName = "postgres"
	return nil
}
