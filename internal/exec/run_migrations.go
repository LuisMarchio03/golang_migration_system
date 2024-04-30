package exec

import (
	"context"
	"database/sql"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"os"
	"path/filepath"
)

// RunMigrations executa todas as migrações encontradas no diretório migrationsDir no banco de dados especificado.
// Retorna um possível erro, se houver.
func RunMigrations(db *sql.DB, migrationsDir string) error {
	// 1. Verificar se o diretório de migrações existe
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("O diretório de migrações não existe")
	}

	// 2. Listar arquivos de migração
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("Erro ao listar arquivos de migração: %v", err)
	}

	// 3. Executar migrações
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".sql" {
			migrationPath := filepath.Join(migrationsDir, f.Name())
			fmt.Println("Executando migração:", migrationPath)

			// Lê o conteúdo do arquivo de migração
			query, err := ioutil.ReadFile(migrationPath)
			if err != nil {
				return fmt.Errorf("Erro ao ler arquivo de migração %s: %v", migrationPath, err)
			}

			// Executa a migração
			_, err = db.Exec(string(query))
			if err != nil {
				return fmt.Errorf("Erro ao executar migração %s: %v", migrationPath, err)
			}

			fmt.Println("Migração concluída com sucesso.")
		}
	}

	return nil
}

// RunMigrationsMongoDB executa todas as migrações encontradas no diretório migrationsDir no banco de dados MongoDB especificado.
// Retorna um possível erro, se houver.
func RunMigrationsMongoDB(db *mongo.Database, migrationsDir string) error {
	// Verificar se o diretório de migrações existe
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return fmt.Errorf("O diretório de migrações não existe")
	}

	// Listar arquivos de migração
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("Erro ao listar arquivos de migração: %v", err)
	}

	// Executar migrações
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".js" {
			migrationPath := filepath.Join(migrationsDir, f.Name())
			fmt.Println("Executando migração:", migrationPath)

			// Ler o conteúdo do arquivo de migração
			scriptBytes, err := ioutil.ReadFile(migrationPath)
			if err != nil {
				return fmt.Errorf("Erro ao ler arquivo de migração %s: %v", migrationPath, err)
			}

			// Executar o script de migração no MongoDB
			result := db.RunCommand(context.Background(), bson.D{
				{"eval", string(scriptBytes)},
			})
			if err := result.Err(); err != nil {
				return fmt.Errorf("Erro ao executar migração %s: %v", migrationPath, err)
			}

			fmt.Println("Migração concluída com sucesso.")
		}
	}

	return nil
}
