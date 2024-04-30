package main

import (
	"database/sql"
	"fmt"
	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/LuisMarchio03/golang_migration_system/internal/exec"
	"go.mongodb.org/mongo-driver/mongo"
)

// MigrationsDir é o diretório onde as migrações serão geradas e executadas
var migrationsDir string // Variável global para armazenar o diretório de migrações

// SetMigrationsDir configura o diretório onde as migrações serão geradas e executadas
func SetMigrationsDir(dir string) {
	migrationsDir = dir
}

// GetMigrationsDir retorna o diretório atualmente configurado para as migrações
func GetMigrationsDir() string {
	return migrationsDir
}

// Cfg representa a configuração do banco de dados
type Cfg = config.Cfg

// Schema representa um esquema de tabela
type Schema = config.Schema

// ConfigDB configura e retorna uma conexão com o banco de dados
func ExecConfigDB(dbDriver string, cfg config.Cfg, migrationsDir string) (*sql.DB, *mongo.Database, error) {
	db, dbNoSql, err := exec.ConfigDB(dbDriver, cfg)
	if err != nil {
		return nil, nil, err
	}

	// Define o diretório de migrações
	SetMigrationsDir(migrationsDir)

	return db, dbNoSql, nil
}

// GenerateMigration gera um arquivo de migração com as schemas fornecidas
func ExecGenerateMigration(schemas ...config.Schema) (string, error) {
	migrationFileName, err := exec.GenerateMigration(migrationsDir, schemas...)
	if err != nil {
		return "", err
	}
	return migrationFileName, nil
}

// RunMigrations executa todas as migrações encontradas no diretório especificado
func ExecRunMigrations(db *sql.DB, migrationsDir string) error {
	err := exec.RunMigrations(db, migrationsDir)
	if err != nil {
		return err
	}
	return nil
}

// ExecGenerateMigrationMongoDB generates a migration file for MongoDB based on the provided documents.
// Returns the name of the created migration file and any possible error.
func ExecGenerateMigrationMongoDB(schemas ...config.Documents) (string, error) {
	// Prepare to collect converted documents
	var convertedDocuments []map[string]interface{}

	// Iterate over each schema of documents
	for _, schema := range schemas {
		// Append each document from the schema into the collection
		for _, doc := range schema.Documents {
			convertedDocuments = append(convertedDocuments, doc)
		}
	}

	// Call GenerateMigrationMongoDB with the converted documents
	migrationFileName, err := exec.GenerateMigrationMongoDB(migrationsDir, convertedDocuments...)
	if err != nil {
		return "", err
	}

	return migrationFileName, nil
}

// ExecRunMigrationsMongoDB executa todas as migrações encontradas no diretório especificado no MongoDB.
// Retorna um possível erro, se houver.
func ExecRunMigrationsMongoDB(db *mongo.Database, migrationsDir string) error {
	err := exec.RunMigrationsMongoDB(db, migrationsDir)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Configuração do banco de dados MySQL
	db, _, err := ExecConfigDB("mysql", config.Cfg{}, "migrationsDir")
	if err != nil {
		fmt.Println("Error configuring MySQL database:", err)
		return
	}

	// Exemplo: gerar migrações para MySQL
	mysqlSchema := config.Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":         "INT NOT NULL AUTO_INCREMENT PRIMARY KEY",
			"username":   "VARCHAR(50) NOT NULL",
			"email":      "VARCHAR(100) NOT NULL",
			"created_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
			"updated_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		},
	}
	mysqlMigrationFileName, err := ExecGenerateMigration(mysqlSchema)
	if err != nil {
		fmt.Println("Error generating MySQL migration:", err)
		return
	}
	fmt.Println("MySQL migration generated successfully:", mysqlMigrationFileName)

	// Exemplo: executar migrações para MySQL
	err = ExecRunMigrations(db, "migrationsDir")
	if err != nil {
		fmt.Println("Error executing MySQL migrations:", err)
		return
	}
	fmt.Println("MySQL migrations completed successfully.")

	// Configuração do banco de dados MongoDB
	_, dbNoSQL, err := ExecConfigDB("mongodb", config.Cfg{}, "migrationsDir")
	if err != nil {
		fmt.Println("Error configuring MongoDB:", err)
		return
	}

	// Exemplo: gerar migrações para MongoDB
	mongoDocuments := config.Documents{
		CollectionName: "users",
		Documents: []map[string]interface{}{
			{"name": "John", "age": 30, "city": "New York"},
			{"name": "Alice", "age": 25, "city": "San Francisco"},
		},
	}
	mongoMigrationFileName, err := ExecGenerateMigrationMongoDB(mongoDocuments)
	if err != nil {
		fmt.Println("Error generating MongoDB migration:", err)
		return
	}
	fmt.Println("MongoDB migration generated successfully:", mongoMigrationFileName)

	// Exemplo: executar migrações para MongoDB
	err = ExecRunMigrationsMongoDB(dbNoSQL, "migrationsDir")
	if err != nil {
		fmt.Println("Error executing MongoDB migrations:", err)
		return
	}
	fmt.Println("MongoDB migrations completed successfully.")
}
