package db

import (
	_ "github.com/go-sql-driver/mysql"
	"hello/pkg/db/connect"
)


func GetConnectAdapter(typeName string) connect.Connecter {
	return connect.NewGorm()
}