package drivers

import (
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/go-sql-driver/mysql"
)

// DbMysql estabelece uma conexão com um banco de dados MySQL utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como usuário, senha, endereço e nome do banco de dados.
// Retorna um ponteiro para sql.DB, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	cfg := config.Cfg{
//	    User:   "username",
//	    Passwd: "password",
//	    Net:    "tcp",
//	    Addr:   "localhost:3306",
//	    DBName: "database_name",
//	}
//	db, err := drivers.DbMysql(cfg)
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//	defer db.Close()
//
//	- Agora você pode usar 'db' para realizar operações no banco de dados MySQL.
//
// Observações:
//   - Certifique-se de que o driver MySQL para Go esteja instalado e disponível no ambiente.
//   - A porta padrão para o MySQL é 3306. Altere o endereço e a porta conforme necessário.
//   - Certifique-se de que o banco de dados MySQL esteja em execução e acessível no endereço especificado.
//   - O usuário e a senha devem ser fornecidos de acordo com as configurações de segurança do seu banco de dados.
func DbMysql(cfg config.Cfg) (*sql.DB, error) {
	cfgMysql := mysql.Config{
		User:   cfg.User,
		Passwd: cfg.Passwd,
		Net:    cfg.Net,
		Addr:   cfg.Addr,
		DBName: cfg.DBName,
	}

	db, err := sql.Open("mysql", cfgMysql.FormatDSN())
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	return db, nil
}
