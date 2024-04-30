package drivers

import (
	"context"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbMongoDB estabelece uma conexão com um banco de dados MongoDB utilizando as configurações fornecidas.
// Recebe um struct Cfg contendo os detalhes de configuração do banco de dados, como endereço e nome do banco de dados.
// Retorna um ponteiro para mongo.Database, que representa a conexão com o banco de dados, e um possível erro, se houver.
//
// Exemplo de uso:
//
//	cfg := config.Cfg{
//	    Addr:   "localhost:27017",
//	    DBName: "database_name",
//	}
//	db, err := drivers.DbMongoDB(cfg)
//	if err != nil {
//	    log.Fatal("Erro ao conectar ao banco de dados:", err)
//	}
//
//	- Agora você pode usar 'db' para realizar operações no banco de dados MongoDB.
func DbMongoDB(cfg config.Cfg) (*mongo.Database, *mongo.Client, error) {
	// Configura as opções para a conexão com o banco de dados MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.Addr)

	// Cria um novo cliente MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, nil, err
	}

	// Verifica se a conexão com o cliente MongoDB é bem-sucedida
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, nil, err
	}

	// Obtém o banco de dados especificado nas configurações
	db := client.Database(cfg.DBName)

	return db, client, nil
}
