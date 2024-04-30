package main_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	golang_migration_system "github.com/LuisMarchio03/golang_migration_system"
	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestExecConfigDB(t *testing.T) {
	// Configuração dos parâmetros do teste
	dbDriver := "mysql"
	cfg := config.Cfg{
		User:   "meu_app_user",
		Passwd: "meu_app_password",
		Net:    "tcp",
		Addr:   "localhost:3308",
		DBName: "meu_app_db",
	}
	migrationsDir := "."

	// Executa a função a ser testada
	db, _, err := golang_migration_system.ExecConfigDB(dbDriver, cfg, migrationsDir)

	// Verifica se não houve erro na configuração do banco de dados
	assert.NoError(t, err, "Erro ao configurar o banco de dados")

	// Verifica se a conexão com o banco de dados foi criada corretamente
	assert.NotNil(t, db, "Conexão com o banco de dados não está definida")

	// Fecha a conexão com o banco de dados no final do teste
	defer db.Close()
}

func TestExecConfigMongoDB(t *testing.T) {
	// Configuração dos parâmetros do teste
	dbDriver := "mongodb"
	cfg := config.Cfg{
		Addr:   "localhost:27017",
		DBName: "mydatabase",
	}
	migrationsDir := "."

	// Executa a função a ser testada
	ctx := context.Background() // Contexto de fundo
	_, dbNoSql, err := golang_migration_system.ExecConfigDB(dbDriver, cfg, migrationsDir)

	// Verifica se não houve erro na configuração do banco de dados
	assert.NoError(t, err, "Erro ao configurar o banco de dados")

	// Verifica se a conexão com o banco de dados foi criada corretamente
	assert.NotNil(t, dbNoSql, "Conexão com o banco de dados NoSQL não está definida")

	// Fecha a conexão com o banco de dados no final do teste
	defer func() {
		err := dbNoSql.Client().Disconnect(ctx)
		assert.NoError(t, err, "Erro ao fechar a conexão com o banco de dados NoSQL")
	}()
}

func TestExecGenerateMigration(t *testing.T) {
	// Cria um exemplo de schema para teste
	schema := config.Schema{
		DbType:    "mysql",
		TableName: "test",
		Fields: map[string]string{
			"id":       "INT AUTO_INCREMENT PRIMARY KEY",
			"username": "VARCHAR(50)",
			"email":    "VARCHAR(100)",
		},
	}

	// Executa a função a ser testada
	migrationFileName, err := golang_migration_system.ExecGenerateMigration(schema)

	// Verifica se não houve erro na geração da migração
	assert.NoError(t, err, "Erro ao gerar a migração")

	// Verifica se o nome do arquivo de migração foi retornado corretamente
	assert.NotEmpty(t, migrationFileName, "Nome do arquivo de migração não está definido")
}

func TestExecRunMigrations(t *testing.T) {
	cfg := config.Cfg{
		User:   "meu_app_user",
		Passwd: "meu_app_password",
		Net:    "tcp",
		Addr:   "localhost:3308",
		DBName: "meu_app_db",
	}
	connString := fmt.Sprintf("%s:%s@%s(%s)/%s",
		cfg.User,
		cfg.Passwd,
		cfg.Net,
		cfg.Addr,
		cfg.DBName,
	)

	// Simula a conexão com o banco de dados (pode ser feita com um banco de dados de teste)
	db, _ := sql.Open("mysql", connString)
	defer db.Close()

	// Define o diretório de migrações para teste
	migrationsDir := "."

	// Executa a função a ser testada
	err := golang_migration_system.ExecRunMigrations(db, migrationsDir)

	// Verifica se não houve erro na execução das migrações
	assert.NoError(t, err, "Erro ao executar as migrações")
}
