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
	Post         string
	UserName     string
	Password     string
	DataBaseName string
}

func (p *PostgreSQL) Init() bool {
	// postgres://UserName:Password@Host:Post/DataBaseName
	PostgreSQLInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.Config.UserName,
		p.Config.Password,
		p.Config.Host,
		p.Config.Post,
		p.Config.DataBaseName)
	DataBaseConnect, err := sql.Open("postgres", PostgreSQLInfo)
	checkError(err)

	if recover() != nil {
		return false
	} else {
		p.ConnectObject = DataBaseConnect
		return true
	}
}

func InitDataBaseConnect(Host string, Port string, UserName string, Password string, DataBaseName string) PostgreSQL {
	DataBaseConnectConfig := Config{
		Host:         Host,
		Post:         Port,
		UserName:     UserName,
		Password:     Password,
		DataBaseName: DataBaseName,
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
