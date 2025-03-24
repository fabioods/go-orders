package config

import (
	"os"
)

type Config struct {
	S3Config  S3Config
	RdsClient DBConfig
}

type S3Config struct {
	S3Bucket    string
	S3Region    string
	S3AccessKey string
	S3SecretKey string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() *Config {
	return &Config{
		S3Config:  loadS3Config(),
		RdsClient: loadRdsConfig(),
	}
}

func loadS3Config() S3Config {
	return S3Config{
		S3Bucket:    os.Getenv("S3_BUCKET"),
		S3Region:    os.Getenv("S3_REGION"),
		S3AccessKey: os.Getenv("S3_ACCESS_KEY"),
		S3SecretKey: os.Getenv("S3_SECRET_KEY"),
	}
}

func loadRdsConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("RDS_HOST"),
		Port:     os.Getenv("RDS_PORT"),
		User:     os.Getenv("RDS_ADMIN_NAME"),
		Password: os.Getenv("RDS_ADMIN_PASSWORD"),
		DBName:   os.Getenv("RDS_DB_NAME"),
		SSLMode:  "disable", // ou pegue via env também
	}
}
