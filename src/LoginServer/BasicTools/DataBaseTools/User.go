package DataBaseTools

type User struct {
	DataBaseObject PostgreSQL
}

func (receiver *User) CheckExist() {
	// PostgreSQLObject := InitDataBaseConnect()
}

func (receiver *User) Test() PostgreSQL {
	var config Config
	config = Config{Host: "localhost", Port: "5432", UserName: "postgres", Password: "PostgreSQLTest", DataBaseName: "postgres"}
	postgresql := InitDataBaseConnect(config)

	return postgresql
}
