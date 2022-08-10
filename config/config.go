package config

type Config struct {
	AppPort    string `json:"appport"`
	DBUser     string `json:"dbname"`
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
	c.DBPassword = "p1"
	c.DBUser = "u1"
	return nil
}
