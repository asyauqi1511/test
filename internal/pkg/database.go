package pkg

import (
	"fmt"
	"github.com/asyauqi1511/test/internal/entity"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	// dataSourceFormat is format of options for connect to db.
	dataSourceFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	// driverName is kind of sql we will use.
	driverName = "pgx"
)

func ConnectDB(config entity.DBConfig) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(dataSourceFormat, config.Hostname, config.Port, config.Username, config.Password, config.DatabaseName)

	// Connect to database.
	db, err := sqlx.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}

	// Connection test.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
