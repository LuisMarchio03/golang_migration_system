package drivers

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/LuisMarchio03/golang_migration_system/internal/config"
// 	_ "github.com/mattn/go-oci8"
// )

// // DbOracle estabelece uma conexão com um banco de dados Oracle utilizando as configurações fornecidas.
// // Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como usuário, senha, endereço e nome do serviço.
// // Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
// //
// // Exemplo de uso:
// //
// //	cfg := config.Cfg{
// //	    User:      "username",
// //	    Passwd:    "password",
// //	    Net:       "tcp",
// //	    Addr:      "localhost:1521",
// //	    Service:   "service_name",
// //	}
// //	db, err := drivers.DbOracle(cfg)
// //	if err != nil {
// //	    log.Fatal("Erro ao conectar ao banco de dados:", err)
// //	}
// //	defer db.Close()
// //
// //	- Agora você pode usar 'db' para realizar operações no banco de dados Oracle.
// func DbOracle(cfg config.Cfg) (*sql.DB, error) {
// 	// Monta a string de conexão com o Oracle
// 	connStr := fmt.Sprintf("%s/%s@%s:%s/%s", cfg.User, cfg.Passwd, cfg.Addr, cfg.Port, cfg.Service)

// 	// Abre a conexão com o banco de dados Oracle
// 	db, err := sql.Open("oci8", connStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Verifica se a conexão com o banco de dados é bem-sucedida
// 	err = db.Ping()
// 	if err != nil {
// 		db.Close()
// 		return nil, err
// 	}

// 	return db, nil
// }
