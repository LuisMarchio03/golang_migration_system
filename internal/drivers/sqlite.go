package drivers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DbSQLite estabelece uma conexão com um banco de dados SQLite utilizando o caminho do arquivo do banco de dados fornecido.
// Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	db, err := drivers.DbSQLite("caminho/para/banco_de_dados.db")
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//	defer db.Close()
//
//	- Agora você pode usar 'db' para realizar operações no banco de dados SQLite.
func DbSQLite(dbPath string) (*sql.DB, error) {
	// Abre a conexão com o banco de dados SQLite
	db, err := sql.Open("sqlite3", dbPath)
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
