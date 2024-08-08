package config

type Config struct {
	HouseDB `json:"house_db"`
	UserDB  `json:"user_db"`
	Service `json:"service"`
}

type Service struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type HouseDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	MaxConns int    `json:"max_conns"`
	Sslmode  string `json:"sslmode"`
	User     string `env:"HOUSE-DB_USER,notEmpty"`
	Password string `env:"HOUSE-DB_PASSWORD,notEmpty"`
}
type UserDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	MaxConns int    `json:"max_conns"`
	Sslmode  string `json:"sslmode"`
	User     string `env:"USER-DB_USER,notEmpty"`
	Password string `env:"USER-DB_PASSWORD,notEmpty"`
}
