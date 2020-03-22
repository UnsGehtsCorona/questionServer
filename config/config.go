package config

var instance Config

type Config struct {
	MongoDB MongoDBConfig
}

func Init() {
	instance = DefaultConfig()
	instance.LoadEnv()
}

func DefaultConfig() Config {
	return Config{
		MongoDB: MongoDBConfig{
			Host:     "localhost:27017",
			Username: "root",
			Password: "wirsindcorona",
		},
	}
}

func (c *Config) LoadEnv() {
	c.MongoDB.LoadEnv()
}

func GetConfig() Config {
	return instance
}
