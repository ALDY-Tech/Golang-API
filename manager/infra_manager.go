package manager

import (
	"submission-project-enigma-laundry/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)



type InfraManager interface {
	SqlDB() *sqlx.DB
}

type infraManager struct {
	db *sqlx.DB
	config config.Config
}

func (i *infraManager) SqlDB() *sqlx.DB {
	return i.db
}

func (i *infraManager) initDb(){
    fmt.Println("Connecting to database with DSN:", i.config.DataSourceName)
    db, err := sqlx.Connect("postgres", i.config.DataSourceName)
    if err != nil {
        panic(err)
    }
    i.db = db
}

func NewInfraManager(cfg config.Config) InfraManager {
	infra := infraManager{config : cfg}
	infra.initDb()
	return &infra
}