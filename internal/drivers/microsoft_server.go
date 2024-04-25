package drivers

import (
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	_ "github.com/denisenkom/go-mssqldb"
)

// DbMSSQLServer estabelece uma conexão com um banco de dados Microsoft SQL Server utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como usuário, senha, endereço e nome do banco de dados.
// Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	cfg := config.Cfg{
//	    User:   "username",
//	    Passwd: "password",
//	    Net:    "tcp",
//	    Addr:   "localhost:1433",
//	    DBName: "database_name",
//	}
//	db, err := drivers.DbMSSQLServer(cfg)
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//	defer db.Close()
//
//	- Agora você pode usar 'db' para realizar operações no banco de dados Microsoft SQL Server.
func DbMSSQLServer(cfg config.Cfg) (*sql.DB, error) {
	// Monta a string de conexão com o Microsoft SQL Server
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
		cfg.Addr, cfg.User, cfg.Passwd, cfg.Port, cfg.DBName)

	// Abre a conexão com o banco de dados Microsoft SQL Server
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão com o banco de dados é bem-sucedida
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
