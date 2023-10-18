package config

var DB *gorm.DB

func ConnectDB() {
	connection := fmt.Sprintf("%v:%v@
}