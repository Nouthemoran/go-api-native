package config

type Config struct {
	PORT string
	DB_USER string
	DB_PASS string
	DB_DATABASE string 
	DB_HOST string
	DB_PORT string
}

var ENV Config