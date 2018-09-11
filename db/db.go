package db

import (
	log "github.com/Sirupsen/logrus"

	"github.com/jackc/pgx"
)

func Init() (*pgx.ConnPool, error) {
	connConf, err := pgx.ParseDSN("host=localhost user=postgres")
	if err != nil {
		log.Fatal(err)
	}
	poolConf := pgx.ConnPoolConfig{
		ConnConfig:     connConf,
		MaxConnections: 16,
	}
	return pgx.NewConnPool(poolConf)
}
