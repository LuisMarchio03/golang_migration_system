package main

import (
	"database/sql"

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
