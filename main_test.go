package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigDB_SupportedDriver(t *testing.T) {
	// Definimos um cenário em que o driver é suportado (MySQL)
	cfg := Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	// Chamamos a função ConfigDB com o driver MySQL
	db, err := ConfigDB("MySql", cfg)

	// Verificamos se não há erro retornado e se o db é não nulo
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestConfigDB_UnsupportedDriver(t *testing.T) {
	// Definimos um cenário em que o driver não é suportado
	cfg := Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	// Chamamos a função ConfigDB com um driver não suportado
	db, err := ConfigDB("Postgres", cfg)

	// Verificamos se há um erro retornado e se o db é nulo
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestDbMysql_ValidConfig(t *testing.T) {
	// Configuração válida do banco de dados MySQL
	cfg := Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "testdb",
	}

	// Chamamos a função DbMysql com a configuração válida
	db, err := DbMysql(cfg)

	// Verificamos se não há erro retornado e se o ponteiro db é não nulo
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestDbMysql_InvalidConfig(t *testing.T) {
	// Configuração inválida do banco de dados MySQL
	cfg := Cfg{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		// Nome do banco de dados ausente
		DBName: "",
	}

	// Chamamos a função DbMysql com a configuração inválida
	db, err := DbMysql(cfg)

	// Verificamos se há um erro retornado e se o ponteiro db é nulo
	assert.Error(t, err)
	assert.Nil(t, db)
}

func TestGenerateMigration(t *testing.T) {
	// Definir schemas de exemplo
	schema1 := Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":   "INT NOT NULL AUTO_INCREMENT",
			"name": "VARCHAR(255)",
			"age":  "INT",
		},
	}
	schema2 := Schema{
		TableName: "products",
		Fields: map[string]string{
			"id":    "INT NOT NULL AUTO_INCREMENT",
			"name":  "VARCHAR(255)",
			"price": "DECIMAL(10,2)",
		},
	}

	// Chamar a função GenerateMigration com os schemas de exemplo
	migrationFileName, err := GenerateMigration(schema1, schema2)

	// Verificar se não há erro retornado
	assert.Nil(t, err)

	// Verificar se o nome do arquivo de migração foi gerado corretamente
	assert.NotEmpty(t, migrationFileName)

	// Verificar se o arquivo de migração existe no diretório de migrações
	_, err = os.Stat(filepath.Join(migrationsDir, migrationFileName))
	assert.False(t, os.IsNotExist(err))

	// Ler o conteúdo do arquivo de migração
	fileContent, err := ioutil.ReadFile(filepath.Join(migrationsDir, migrationFileName))
	assert.Nil(t, err)

	// Verificar se o conteúdo do arquivo de migração corresponde ao esperado
	expectedContent := `CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    age INT
);

CREATE TABLE IF NOT EXISTS products (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    price DECIMAL(10,2)
);
`
	assert.Equal(t, expectedContent, string(fileContent))
}

// Generates a migration file with a unique name based on current timestamp
func TestGenerateMigration_UniqueName(t *testing.T) {
	// Create a temporary directory for migrations
	tempDir, err := ioutil.TempDir("", "migrations")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set the migrationsDir to the temporary directory
	migrationsDir = tempDir

	// Call the GenerateMigration function
	_, err = GenerateMigration(Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":       "INT AUTO_INCREMENT PRIMARY KEY",
			"username": "VARCHAR(50)",
			"email":    "VARCHAR(100)",
		},
	})
	if err != nil {
		t.Fatalf("Failed to generate migration: %v", err)
	}

	// Check if the migration file exists in the temporary directory
	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	found := false
	for _, file := range files {
		if file.Name() == "migration_" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Migration file not found")
	}
}

// Returns an error if unable to create the migration file
func TestGenerateMigration_ErrorCreatingFile(t *testing.T) {
	// Create a read-only directory for migrations
	tempDir, err := ioutil.TempDir("", "migrations")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set the migrationsDir to the read-only directory
	migrationsDir = tempDir

	// Call the GenerateMigration function
	_, err = GenerateMigration(Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":       "INT AUTO_INCREMENT PRIMARY KEY",
			"username": "VARCHAR(50)",
			"email":    "VARCHAR(100)",
		},
	})

	// Check if the error is not nil
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}
}

// Executes all migrations in the migrations directory successfully
func TestRunMigrations_SuccessfulExecution(t *testing.T) {
	// Create a temporary migrations directory
	tempDir, err := ioutil.TempDir("", "migrations")
	if err != nil {
		t.Fatalf("Failed to create temporary migrations directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a temporary migration file
	migrationFile := filepath.Join(tempDir, "migration_20220101120000.sql")
	file, err := os.Create(migrationFile)
	if err != nil {
		t.Fatalf("Failed to create temporary migration file: %v", err)
	}
	defer file.Close()

	// Write migration content to the file
	migrationContent := "CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(50), email VARCHAR(100));"
	_, err = file.WriteString(migrationContent)
	if err != nil {
		t.Fatalf("Failed to write migration content to file: %v", err)
	}

	// Create a mock sql.DB
	mockDB := &sql.DB{}

	// Set expectations for db.Exec
	expectedQuery := migrationContent
	mockDB.ExecFunc = func(query string) (sql.Result, error) {
		if query != expectedQuery {
			return nil, fmt.Errorf("Unexpected query: %s", query)
		}
		return nil, nil
	}

	// Call the RunMigrations function
	err = RunMigrations(mockDB)
	if err != nil {
		t.Fatalf("RunMigrations returned an error: %v", err)
	}
}

// Handles non-existent migrations directory gracefully
func TestRunMigrations_NonExistentDirectory(t *testing.T) {
	// Create a mock sql.DB
	mockDB := &sql.DB{}

	// Call the RunMigrations function
	err := RunMigrations(mockDB)
	if err == nil {
		t.Fatal("RunMigrations did not return an error for non-existent migrations directory")
	}
	expectedError := "O diretório de migrações não existe"
	if err.Error() != expectedError {
		t.Fatalf("Unexpected error message. Expected: %s, Got: %s", expectedError, err.Error())
	}
}
