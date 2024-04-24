package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	main "github.com/LuisMarchio03/golang_migration_system"
)

func TestConfigDB_MySQL(t *testing.T) {
	cfg := main.Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	db, err := main.ConfigDB("MySql", cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if db == nil {
		t.Error("expected non-nil db, got nil")
	}
}

func TestConfigDB_UnsupportedDriver(t *testing.T) {
	cfg := main.Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	db, err := main.ConfigDB("Postgres", cfg)

	if err == nil {
		t.Error("expected error for unsupported driver, got nil")
	}
	if db != nil {
		t.Error("expected nil db for unsupported driver, got non-nil")
	}
}

func TestDbMysql(t *testing.T) {
	// Configuração do banco de dados MySQL para o teste
	cfg := main.Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	// Chame a função DbMysql com a configuração fornecida
	db, err := main.DbMysql(cfg)

	// Verifique se não há erro retornado
	if err != nil {
		t.Errorf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Verifique se o objeto de banco de dados não é nulo
	if db == nil {
		t.Error("O objeto de banco de dados é nulo")
	}
}

func TestGenerateMigration(t *testing.T) {
	// Defina um diretório temporário para os testes
	tempDir, err := ioutil.TempDir("", "test_migrations")
	if err != nil {
		t.Fatalf("Erro ao criar diretório temporário: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Defina o diretório de migrações para o diretório temporário
	main.MigrationsDir = tempDir

	// Defina os esquemas de exemplo
	schema1 := main.Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":   "INT NOT NULL AUTO_INCREMENT",
			"name": "VARCHAR(255)",
			"age":  "INT",
		},
	}
	schema2 := main.Schema{
		TableName: "products",
		Fields: map[string]string{
			"id":    "INT NOT NULL AUTO_INCREMENT",
			"name":  "VARCHAR(255)",
			"price": "DECIMAL(10,2)",
		},
	}

	// Chame a função GenerateMigration com os esquemas de exemplo
	migrationFileName, err := main.GenerateMigration(schema1, schema2)

	// Verifique se não há erro retornado
	if err != nil {
		t.Fatalf("Erro ao gerar migração: %v", err)
	}

	// Verifique se o nome do arquivo de migração não está vazio
	if migrationFileName == "" {
		t.Fatal("O nome do arquivo de migração está vazio")
	}

	// Verifique se o arquivo de migração foi criado no diretório de migrações
	_, err = os.Stat(filepath.Join(tempDir, migrationFileName))
	if err != nil {
		t.Fatalf("Erro ao verificar se o arquivo de migração foi criado: %v", err)
	}
}

func TestRunMigrations(t *testing.T) {
	// Configuração do banco de dados MySQL para o teste
	cfg := main.Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	// Chame a função DbMysql com a configuração fornecida
	db, err := main.DbMysql(cfg)
	if err != nil {
		t.Fatalf("Erro ao abrir a conexão com o banco de dados MySQL: %v", err)
	}
	defer db.Close()

	// Criar um diretório temporário para as migrações
	tempDir, err := ioutil.TempDir("", "test_migrations")
	if err != nil {
		t.Fatalf("Erro ao criar o diretório temporário: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Defina o diretório de migrações para o diretório temporário
	main.MigrationsDir = tempDir

	// Criar arquivos de migração temporários dentro do diretório temporário
	migrationFiles := []string{"migration_1.sql", "migration_2.sql"}
	for _, fileName := range migrationFiles {
		filePath := filepath.Join(tempDir, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Erro ao criar o arquivo de migração: %v", err)
		}
		defer file.Close()

		// Escrever conteúdo de exemplo nos arquivos de migração
		_, err = file.WriteString("CREATE TABLE IF NOT EXISTS test_table (id INT);")
		if err != nil {
			t.Fatalf("Erro ao escrever conteúdo no arquivo de migração: %v", err)
		}
	}

	// Executar migrações
	err = main.RunMigrations(db)
	if err != nil {
		t.Fatalf("Erro ao executar migrações: %v", err)
	}

	// Verificar se as tabelas foram criadas no banco de dados
	var tableName string
	err = db.QueryRow("SHOW TABLES LIKE 'test_table'").Scan(&tableName)
	if err != nil {
		t.Fatalf("Erro ao verificar se a tabela foi criada: %v", err)
	}

	// Verificar se o nome da tabela corresponde ao esperado
	expectedTableName := "test_table"
	if tableName != expectedTableName {
		t.Errorf("Nome da tabela incorreto. Esperado: %s, Obtido: %s", expectedTableName, tableName)
	}
}
