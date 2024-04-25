package drivers

import (
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	_ "github.com/lib/pq"
)

// DbPostgreSQL estabelece uma conexão com um banco de dados PostgreSQL utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como usuário, senha, endereço e nome do banco de dados.
// Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	cfg := config.Cfg{
//	    User:   "username",
//	    Passwd: "password",
//	    Net:    "tcp",
//	    Addr:   "localhost:5432",
//	    DBName: "database_name",
//	}
//	db, err := drivers.DbPostgreSQL(cfg)
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//	defer db.Close()
//
//	- Agora você pode usar 'db' para realizar operações no banco de dados PostgreSQL.
func DbPostgreSQL(cfg config.Cfg) (*sql.DB, error) {
	// Monta a string de conexão com o PostgreSQL
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Passwd, cfg.Addr, cfg.Port, cfg.DBName)

	// Abre a conexão com o banco de dados PostgreSQL
	db, err := sql.Open("postgres", connStr)
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
