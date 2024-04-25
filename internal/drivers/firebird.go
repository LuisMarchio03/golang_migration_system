package drivers

import (
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
)

// DbFirebird estabelece uma conexão com um banco de dados Firebird utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como usuário, senha, endereço e nome do banco de dados.
// Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	cfg := config.Cfg{
//	    User:   "username",
//	    Passwd: "password",
//	    Net:    "tcp",
//	    Addr:   "localhost:3050",
//	    DBName: "database_name",
//	}
//	db, err := drivers.DbFirebird(cfg)
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//	defer db.Close()
//
//	- Agora você pode usar 'db' para realizar operações no banco de dados Firebird.
//
// Observações:
//   - Certifique-se de que o driver Firebird para Go esteja instalado e disponível no ambiente.
//   - A porta padrão para o Firebird é 3050. Altere o endereço e a porta conforme necessário.
//   - Certifique-se de que o banco de dados Firebird esteja em execução e acessível no endereço especificado.
//   - O usuário e a senha devem ser fornecidos de acordo com as configurações de segurança do seu banco de dados.
func DbFirebird(cfg config.Cfg) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@%s/%s", cfg.User, cfg.Passwd, cfg.Addr, cfg.DBName)

	db, err := sql.Open("firebirdsql", connString)
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	return db, nil
}
