package config

var DB *gorm.DB

func ConnectDB() {
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vBandung", ENV.DB_USER, ENV.DB_PASS, ENV
}