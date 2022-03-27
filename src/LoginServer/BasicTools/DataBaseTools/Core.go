package DataBaseTools

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	Config        Config
	ConnectObject *sql.DB
}
type Config struct {
	Host         string
	Port         string
	UserName     string
	Password     string
	DataBaseName string
}

func (p *PostgreSQL) Init() {
	// postgres://UserName:Password@Host:Port/DataBaseName
	PostgreSQLInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.Config.UserName,
		p.Config.Password,
		p.Config.Host,
		p.Config.Port,
		p.Config.DataBaseName)
	DataBaseConnect, err := sql.Open("postgres", PostgreSQLInfo)
	checkError(err)

	p.ConnectObject = DataBaseConnect
}

func InitDataBaseConnect(config Config) PostgreSQL {
	DataBaseConnectConfig := Config{
		Host:         config.Host,
		Port:         config.Port,
		UserName:     config.UserName,
		Password:     config.Password,
		DataBaseName: config.DataBaseName,
	}
	tmpStruct := PostgreSQL{
		Config: DataBaseConnectConfig,
	}

	tmpStruct.Init()
	return tmpStruct
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
