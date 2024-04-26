package exec

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/LuisMarchio03/golang_migration_system/internal/drivers"
)

// ConfigDB configura o banco de dados com base no driver especificado e nas configurações fornecidas.
// Ele recebe o nome do driver do banco de dados e as configurações do banco de dados como parâmetros.
// Se o driver for "MySql", chama a função DbMysql para configurar o banco de dados MySQL.
// Retorna um possível erro, se houver.
func ConfigDB(dbDriver string, cfg config.Cfg) (*sql.DB, error) {
	var db *sql.DB
	var err error

	dbDriver = strings.ToLower(dbDriver)

	switch dbDriver {
	case "mysql":
		db, err = drivers.DbMysql(cfg)
	case "firebirdsql":
		db, err = drivers.DbFirebird(cfg)
	case "postgresql":
		db, err = drivers.DbPostgreSQL(cfg)
	default:
		return nil, fmt.Errorf("Driver de banco de dados não suportado: %s", dbDriver)
	}

	return db, err
}
