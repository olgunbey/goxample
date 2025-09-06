package config

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func GetConfig() *Postgres {

	return &Postgres{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "admin",
		Dbname:   "goxample",
	}
}
