package configs

var Port *int
var DbConfig *DBConfig
var JwtSecret *string

func BootConfig() {
	env := NewEnvironment()

	JwtSecret = env.GetString("JWT_SECRET")
	DbConfig = &DBConfig{
		Host:     *env.GetString("DB_HOST"),
		Port:     *env.GetString("DB_PORT"),
		Username: *env.GetString("DB_USER"),
		Password: *env.GetString("DB_PASSWORD"),
		DBName:   *env.GetString("DB_NAME"),
	}
}
