package DataBaseTools

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Connect interface {
	Init() bool
}

type PostgreSQL struct {
	Host          string
	Post          string
	UserName      string
	Password      string
	DataBaseName  string
	ConnectObject *sql.DB
}

func (p PostgreSQL) Init() bool {
	// postgres://UserName:Password@Host:Post/DataBaseName
	PostgreSQLInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.UserName,
		p.Password,
		p.Host,
		p.Post,
		p.DataBaseName)
	DataBaseConnect, err := sql.Open("postgres", PostgreSQLInfo)
	checkError(err)

	if recover() != nil {
		return false
	} else {
		p.ConnectObject = DataBaseConnect
		return true
	}
}

func InitDataBaseConnect(Host string, Port string, UserName string, Password string, DataBaseName string) Connect {
	var tmpInterface Connect
	tmpInterface = PostgreSQL{
		Host:         Host,
		Post:         Port,
		UserName:     UserName,
		Password:     Password,
		DataBaseName: DataBaseName,
	}
	tmpInterface.Init()
	return tmpInterface
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
