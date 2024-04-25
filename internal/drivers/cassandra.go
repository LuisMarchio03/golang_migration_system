package drivers

import (
	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/gocql/gocql"
)

// DbCassandra estabelece uma conexão com um banco de dados Cassandra utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como endereço e nome do keyspace.
// Retorna uma sessão de cassandra.Session, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	cfg := config.Cfg{
//	    Addr:     "localhost:9042",
//	    Keyspace: "keyspace_name",
//	}
//	session, err := drivers.DbCassandra(cfg)
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//	defer session.Close()
//
//	- Agora você pode usar 'session' para realizar operações no banco de dados Cassandra.
func DbCassandra(cfg config.Cfg) (*gocql.Session, error) {
	// Cria a configuração para a conexão com o banco de dados Cassandra
	cluster := gocql.NewCluster(cfg.Addr)
	cluster.Keyspace = cfg.Keyspace

	// Conecta ao cluster Cassandra
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	// Retorna a sessão Cassandra
	return session, nil
}
