package initialize

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq" 

	database "github.com/augustus281/trackingcoin/database/sqlc"
	"github.com/augustus281/trackingcoin/global"
)

func InitDB() {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		global.Config.PostgreSql.Username,
		global.Config.PostgreSql.Password,
		global.Config.PostgreSql.Host,
		global.Config.PostgreSql.Port,
		global.Config.PostgreSql.DBName,
		global.Config.PostgreSql.SslMode,
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
