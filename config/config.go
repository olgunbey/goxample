package config

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func GetPostgresConfig() *Postgres {

	return &Postgres{
		Host:     "localhost",
		Port:     5432,
		User:     "user123",
		Password: "password123",
		Dbname:   "GoxampleDb",
	}
}
