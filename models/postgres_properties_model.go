package models

type PostgresProperties struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func (p PostgresProperties) GetDSN() string {
	return "host=" + p.Host + " user=" + p.User + " password=" + p.Password + " dbname=" + p.DBName + " port=" + p.Port + " sslmode=" + p.SSLMode + " TimeZone=" + p.TimeZone
}
