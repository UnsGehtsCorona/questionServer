package config

import "os"

type MongoDBConfig struct {
	Host     string
	Username string
	Password string
}

func (m *MongoDBConfig) LoadEnv() {
	if host := os.Getenv("MONGO_HOST"); host != "" {
		m.Host = host
	}

	if username := os.Getenv("MONGO_USERNAME"); username != "" {
		m.Username = username
	}

	if password := os.Getenv("MONGO_PASSWORD"); password != "" {
		m.Host = password
	}
}
