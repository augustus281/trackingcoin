package initialize

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"

	database "github.com/augustus281/trackingcoin/database/sqlc"
	"github.com/augustus281/trackingcoin/global"
)

func InitDB() {
	username := "root"
	password := global.Config.PostgreSql.Password
	host := "localhost"
	port := 5432
	dbName := global.Config.PostgreSql.DBName
	sslMode := "disable"

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		username,
		password,
		host,
		port,
		dbName,
		sslMode,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect database", err)
	}
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)

	global.Logger.Info("Connect database successfully!")
	global.Db = database.NewStore(conn)
}
