package main

import (
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang_migration_system/internal/config"
	"github.com/LuisMarchio03/golang_migration_system/internal/exec"
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
func ExecConfigDB(dbDriver string, cfg config.Cfg, migrationsDir string) (*sql.DB, error) {
	db, err := exec.ConfigDB(dbDriver, cfg)
	if err != nil {
		return nil, err
	}

	// Define o diretório de migrações
	SetMigrationsDir(migrationsDir)

	return db, nil
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

func main() {
	// Aqui você pode chamar suas funções conforme necessário
	db, err := ExecConfigDB("mysql", config.Cfg{}, "migrationsDir")
	if err != nil {
		fmt.Println("Error configuring database:", err)
		return
	}

	// Exemplo: gerar migrações
	schema := config.Schema{
		TableName: "users",
		Fields: map[string]string{
			"id":         "INT NOT NULL AUTO_INCREMENT PRIMARY KEY",
			"username":   "VARCHAR(50) NOT NULL",
			"email":      "VARCHAR(100) NOT NULL",
			"created_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
			"updated_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		},
	}
	migrationFileName, err := ExecGenerateMigration(schema)
	if err != nil {
		fmt.Println("Error generating migration:", err)
		return
	}
	fmt.Println("Migration generated successfully:", migrationFileName)

	// Exemplo: executar migrações
	err = ExecRunMigrations(db, "migrationsDir")
	if err != nil {
		fmt.Println("Error executing migrations:", err)
		return
	}
	fmt.Println("Migrations completed successfully.")
}
