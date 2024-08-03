package config

type Config struct {
	Postgres `json:"postgres"`
	Service  `json:"service"`
}

type Service struct {
	Port string `json:"port"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	MaxConns int    `json:"max_conns"`
	Sslmode  string `json:"sslmode"`
	User     string `env:"PG_USER,notEmpty"`
	Password string `env:"PG_PASSWORD,notEmpty"`
}
