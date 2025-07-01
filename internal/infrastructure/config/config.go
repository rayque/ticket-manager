package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppName          string
	MongoDBUri       string
	MongoDBDatabase  string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresPort     string
	PostgresHost     string
	JWTSecret        string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		AppName:          os.Getenv("APP_NAME"),
		MongoDBUri:       os.Getenv("MONGO_URI"),
		MongoDBDatabase:  os.Getenv("MONGO_DATABASE"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		JWTSecret:        getJWTSecret(),
	}
}

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("JWT_SECRET não encontrado no .env, usando valor padrão (NÃO RECOMENDADO PARA PRODUÇÃO)")
		return "default-jwt-secret-key-change-in-production"
	}
	return secret
}
